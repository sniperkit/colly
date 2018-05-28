package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/countlog"
	"sync"
	"context"
)

// node.go use topo.go to know the master/slaves of partitionId
// topo.go use node.go to know the nodes and update topo to optimal state
// node.go will trigger topo.go to update topo if nodes changed

type partitionNodes struct {
	isPromoting bool
	master string
	slaves []string
}

type topology []partitionNodes

var topo = make(topology, kvstore.PartitionsCount)
var topoMutex = &sync.Mutex{}

func (topo topology) get(partitionId uint64) partitionNodes {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	return topo[partitionId]
}

func (topo topology) setPromotingMaster(partitionId uint64, master string) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	topo[partitionId].master = master
	topo[partitionId].isPromoting = true
	if thisNodeStarted {
		countlog.Info("event!topo.promoting master",
			"partitionId", partitionId,
			"master", master)
	}
}

func (topo topology) setPromotedMaster(partitionId uint64, master string) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	topo[partitionId].master = master
	topo[partitionId].isPromoting = false
	countlog.Info("event!topo.promoted master",
		"partitionId", partitionId,
			"master", master)
}

// if the master changed, clear the master
// it will be filled when rebalance
func (topo topology) clearMaster(partitionId uint64) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	oldMaster := topo[partitionId].master
	topo[partitionId].master = ""
	topo[partitionId].isPromoting = false
	commandProcessorsMutexs[partitionId].Lock()
	commandProcessors[partitionId] = nil
	commandProcessorsMutexs[partitionId].Unlock()
	countlog.Info("event!topo.clear master",
		"partitionId", partitionId,
		"oldMaster", oldMaster)
	triggerTopoChanged()
}

func (topo topology) refresh(ctx context.Context, partitionId uint64) (string, error) {
	metadataKey := fmt.Sprintf("partition_%v", partitionId)
	encodedPartition, err := kvstore.GetMetadata(ctx, metadataKey)
	if err != nil {
		countlog.Error("event!topo.failed to get partitionId topo from kvstore",
			"err", err,
			"metadataKey", metadataKey,
			"partitionId", partitionId)
		return "", err
	}
	if len(encodedPartition) == 0 {
		return "", nil
	}
	master := jsoniter.Get(encodedPartition, "Master").ToString()
	if master != "" {
		topo.setPromotedMaster(partitionId, master)
	}
	return master, nil
}

// rebalance is only triggered from node.go
// given this cluster and your current topo
// update the topo to a new optimal state
func (topo topology) rebalance(ctx context.Context, cluster map[string]*nodeStatus) error {
	var err error
	myMasterPartitionsCount := 0
	for partitionId := uint64(0); partitionId < kvstore.PartitionsCount; partitionId++ {
		master := topo.get(partitionId).master
		if master == "" {
			master, err = topo.refresh(ctx, partitionId)
			if err != nil {
				return err
			}
		}
		if master == "" {
			node := chooseMaster(cluster)
			// will be reset when refreshing from kvstore
			// only used in this rebalancing cycle
			node.MasterPartitionsCount += 1
			topo.setPromotingMaster(partitionId, node.Addr)
		}
		if master == thisNodeAddr {
			myMasterPartitionsCount += 1
		}
	}
	cluster[thisNodeAddr].MasterPartitionsCount = myMasterPartitionsCount
	return nil
}

func chooseMaster(cluster map[string]*nodeStatus) *nodeStatus {
	var chosen *nodeStatus
	for _, node := range cluster {
		if chosen == nil {
			chosen = node
		} else {
			if node.MasterPartitionsCount < chosen.MasterPartitionsCount {
				chosen = node
			}
		}
	}
	return chosen
}

func (topo topology) savePromotedMasterInBackground(ctx context.Context, partitionId uint64) {
	master := thisNodeAddr
	topo.setPromotedMaster(partitionId, master)
	encodedPartition, err := jsoniter.Marshal(map[string]string{"Master":master})
	if err != nil {
		countlog.Error("event!topo.failed to marshal partitionId topo",
			"partitionId", partitionId, "err", err)
		return
	}
	metadataKey := fmt.Sprintf("partition_%v", partitionId)
	err = kvstore.SetMetadata(ctx, metadataKey, encodedPartition)
	if err != nil {
		countlog.Error("event!topo.failed to save partitionId topo",
			"partitionId", partitionId, "err", err)
		return
	}
	// TODO: broadcast new master to the cluster
}
