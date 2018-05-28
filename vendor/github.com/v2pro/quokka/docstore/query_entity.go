package docstore

import (
	"net/http"
	"strings"
	"github.com/v2pro/quokka/kvstore"
	"github.com/v2pro/plz/countlog"
	"github.com/hashicorp/golang-lru"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
)

type entityCache struct {
	cache         *lru.TwoQueueCache
	kvstoreLookup *kvstoreLookup
}

func newEntityCache() *entityCache {
	cache, _ := lru.New2Q(1024 * 32)
	return &entityCache{
		cache: cache,
		kvstoreLookup: &kvstoreLookup{"entity_"},
	}
}

func (cache *entityCache) queryEntity(respWriter http.ResponseWriter, req *http.Request) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!queryEntity.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	path := strings.TrimRight(req.URL.Path, "/")
	parts := strings.Split(path, "/")
	entityId := parts[len(parts)-1]
	entityType := parts[len(parts)-3]
	partitionId := kvstore.HashToPartition(entityId)
	entityObj, _ := cache.cache.Get(entityId)
	entity, _ := entityObj.(*entity)
	if entity == nil {
		entity = newEntity()
	}
	eventId, err := cache.kvstoreLookup.getEventId(req.Context(), partitionId, entityType, entityId, 0)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	if eventId == 0 {
		respWriter.Write(replyError(withErrorNumber(errors.New("entity missing"), ErrEntityMissing)))
		return
	}
	err = loadEntity(req.Context(), partitionId, entityType, entityId, eventId, entity)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	cache.cache.Add(entityId, entity)
	stream := runtime.Json.BorrowStream(nil)
	defer runtime.Json.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	stream.WriteVal(0)
	stream.WriteMore()
	stream.WriteObjectField("entity_version")
	stream.WriteVal(entity.version)
	stream.WriteMore()
	stream.WriteObjectField("data")
	stream.WriteVal(entity.doc)
	stream.WriteObjectEnd()
	if stream.Error != nil {
		respWriter.Write(replyError(err))
		return
	}
	respWriter.Write(stream.Buffer())
	return
}
