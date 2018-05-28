package memkv

import (
	"github.com/v2pro/quokka/kvstore"
	"errors"
	"context"
)

type memRow []byte
type memPartition [][]byte

var memPartitions = make([]memPartition, kvstore.PartitionsCount)
var memMetadata = []*kvstore.MetadataRow{}
var memMonotonics = []map[string]uint64{}

func ResetKVStore() {
	memPartitions = make([]memPartition, kvstore.PartitionsCount)
	memMonotonics = make([]map[string]uint64, kvstore.PartitionsCount)
	for i := 0; i < len(memMonotonics); i++ {
		memMonotonics[i] = map[string]uint64{}
	}
	memMetadata = []*kvstore.MetadataRow{}
	kvstore.Get = memGet
	kvstore.Append = memAppend
	kvstore.Scan = memScan
	kvstore.GetMetadata = memGetMetadata
	kvstore.SetMetadata = memSetMetadata
	kvstore.ScanMetadata = memScanMetadata
	kvstore.GetMonotonic = memGetMonotonic
	kvstore.SetMonotonic = memSetMonotonic
}

func memGet(ctx context.Context, partition uint64, namespace string, rowKey uint64) ([]byte, error) {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return nil, errors.New("partition not found")
	}
	rows := memPartitions[partition]
	if rowKey <= 0 || rowKey >= uint64(len(rows)) {
		return nil, errors.New("row not found")
	}
	return rows[rowKey], nil
}

func memAppend(ctx context.Context, partition uint64, namespace string, rowKey uint64, rowValue []byte) error {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return errors.New("partition not found")
	}
	rows := memPartitions[partition]
	if rows == nil {
		rows = make(memPartition, 1)
	}
	if rowKey < uint64(len(rows)) {
		return kvstore.KeyConflictError
	}
	if rowKey != uint64(len(rows)) {
		return errors.New("row key should be +1")
	}
	rows = append(rows, rowValue)
	memPartitions[partition] = rows
	return nil
}

func memScan(ctx context.Context, partition uint64, namespace string, fromKey uint64) (kvstore.RowIterator, error) {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return nil, errors.New("partition not found")
	}
	rows := memPartitions[partition]
	currentOffset := fromKey
	return func() ([]kvstore.Row, error) {
		if currentOffset >= uint64(len(rows)) {
			return nil, nil
		}
		nextOffset := currentOffset + 100
		if nextOffset > uint64(len(rows)) {
			nextOffset = uint64(len(rows))
		}
		batchValues := rows[currentOffset:nextOffset]
		batchRows := make([]kvstore.Row, len(batchValues))
		for i := 0; i < len(batchValues); i++ {
			batchRows[i].Key = currentOffset + uint64(i)
			batchRows[i].Value = batchValues[i]
		}
		currentOffset = nextOffset
		return batchRows, nil
	}, nil
}

func memGetMonotonic(ctx context.Context, partition uint64, namespace string, key string) (uint64, error) {
	return memMonotonics[partition][key], nil
}

func memSetMonotonic(ctx context.Context, partition uint64, namespace string, key string, value uint64) error {
	if value > memMonotonics[partition][key] {
		memMonotonics[partition][key] = value
	}
	return nil
}

func memGetMetadata(ctx context.Context, key string) ([]byte, error) {
	for _, row := range memMetadata {
		if row.Key == key {
			return row.Value, nil
		}
	}
	return nil, nil
}

func memSetMetadata(ctx context.Context, key string, value []byte) error {
	for _, row := range memMetadata {
		if row.Key == key {
			row.Value = value
			return nil
		}
	}
	memMetadata = append(memMetadata, &kvstore.MetadataRow{
		Key:   key,
		Value: value,
	})
	return nil
}

func memScanMetadata(ctx context.Context, fromKey string, toKey string) (kvstore.MetadataRowIterator, error) {
	returned := false
	return func() ([]kvstore.MetadataRow, error) {
		if returned {
			return nil, nil
		}
		returned = true
		found := []kvstore.MetadataRow{}
		for _, row := range memMetadata {
			if row.Key >= fromKey && row.Key < toKey {
				found = append(found, *row)
			}
		}
		return found, nil
	}, nil
}
