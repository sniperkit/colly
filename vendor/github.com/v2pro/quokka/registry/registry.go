package registry

import (
	"net"
	"time"
)

type ProcessPerceivedState struct {
	WasAlive       bool
	LastHeartbeat  time.Time
	ObjectVersions map[ObjectId]ObjectVersion
}

type ProcessDesiredState struct {
	// only this flag set to true, the process is considered as part of distribution
	ShouldBeAlive bool
}

type Process struct {
	Node           net.IPAddr
	ProcessLocalId string // unique within a node, normally the startup path
	Cluster        string
	PerceivedState *ProcessPerceivedState
	// if no desired state, the process is shown up unexpectedly causing by dynamic registration
	DesiredState *ProcessDesiredState
}

type ClusterId struct {
	ServiceName string // globally unique
	ClusterName string // unique within service
}

type ServiceDesiredState struct {
	Tags map[string]string
}

type Service struct {
	ServiceName string
	// if no desired state, the service is shown up unexpectedly causing by dynamic registration
	DesiredState *ServiceDesiredState
}

type ClusterDesiredState struct {
	Configs []ObjectId
	Tags    map[string]string
}

type Cluster struct {
	ClusterId ClusterId
	// if no desired state, the cluster is shown up unexpectedly causing by dynamic registration
	DesiredState *ClusterDesiredState
}

type ObjectId struct {
	Namespace  string
	ObjectName string
}

type Object struct {
	ObjectId      ObjectId
	Toggle        ObjectItem
	Variants      []ObjectItem
	ObjectVersion ObjectVersion
}

type ObjectVersion int64 // timestamp

type ObjectItem struct {
	ItemHash   string // actual content in binstore
	ItemName   string
	RefObjects []ObjectId
}
