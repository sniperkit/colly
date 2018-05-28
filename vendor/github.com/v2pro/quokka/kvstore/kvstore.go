package kvstore

import (
	"errors"
	"github.com/spaolacci/murmur3"
	"context"
)

const ZonesCount = 32
const PartitionsPerZone = 997
const PartitionsCount = ZonesCount * PartitionsPerZone

var KeyConflictError = errors.New("key conflict")

func HashToPartition(entityId string) uint64 {
	hash := murmur3.Sum64([]byte(entityId[:len(entityId) - 1]))
	// by control the first byte, we can control the partition of the entity
	// for example, we can assign large entity type to
	// not sharing partitions with other entity types
	zoneId := lookupZone(entityId[len(entityId) - 1])
	partition := hash % 997
	return zoneId * partition
}

// kvstore for event log

type Row struct {
	Key   uint64
	Value []byte
}
type RowIterator func() ([]Row, error)

var Get = func(ctx context.Context, partition uint64, namespace string, key uint64) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var Scan = func(ctx context.Context, partition uint64, namespace string, fromKey uint64) (RowIterator, error) {
	return nil, errors.New("not implemented")
}

var Append = func(ctx context.Context, partition uint64, namespace string, key uint64, rowValue []byte) error {
	return errors.New("not implemented")
}

// entity lookup:
// entity id => event id
// command lookup:
// command id => event id

var GetMonotonic = func(ctx context.Context, partition uint64, namespace string, key string) (uint64, error) {
	return 0, errors.New("not implemented")
}

var SetMonotonic = func(ctx context.Context, partition uint64, namespace string, key string, value uint64) error {
	return errors.New("not implemented")
}

// metadata:
// partition master and slaves
// server heartbeats

type MetadataRow struct {
	Key   string
	Value []byte
}
type MetadataRowIterator func() ([]MetadataRow, error)

var GetMetadata = func(ctx context.Context, key string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var SetMetadata = func(ctx context.Context, key string, value []byte) error {
	return errors.New("not implemented")
}

// [fromKey, toKey] ordered lexically
var ScanMetadata = func(ctx context.Context, fromKey string, toKey string) (MetadataRowIterator, error) {
	return nil, errors.New("not implemented")
}
