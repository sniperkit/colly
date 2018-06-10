package tablib

import (
	"bytes"
	"encoding/csv"
	// "sync"
)

// CSV returns a CSV representation of the Dataset an Export.
func (d *Dataset) CSV() (*Export, error) {
	records := d.Records()
	b := newBuffer()

	w := csv.NewWriter(b)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		return nil, err
	}

	return newExport(b), nil
}

// LoadCSV loads a Dataset by its CSV representation.
func LoadCSV(input []byte) (*Dataset, error) {
	reader := csv.NewReader(bytes.NewReader(input))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	ds := NewDataset(records[0])
	for i := 1; i < len(records); i++ {
		// this is odd
		row := make([]interface{}, len(records[i]))
		for k, v := range records[i] {
			row[k] = v
		}
		ds.Append(row)
	}

	return ds, nil
}

// TSV returns a TSV representation of the Dataset as string.
func (d *Dataset) TSV() (*Export, error) {
	records := d.Records()
	b := newBuffer()

	w := csv.NewWriter(b)
	w.Comma = '\t'
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		return nil, err
	}

	return newExport(b), nil
}

// LoadTSV loads a Dataset by its TSV representation.
func LoadTSV(input []byte) (*Dataset, error) {
	reader := csv.NewReader(bytes.NewReader(input))
	reader.Comma = '\t'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	ds := NewDataset(records[0])
	for i := 1; i < len(records); i++ {
		// this is odd
		row := make([]interface{}, len(records[i]))
		for k, v := range records[i] {
			row[k] = v
		}
		ds.Append(row)
	}

	return ds, nil
}

func toRecord(kv KeyValue, keys []string) []string {
	record := make([]string, 0, len(keys))
	for _, key := range keys {
		if value, ok := kv[key]; ok {
			record = append(record, toString(value))
		} else {
			record = append(record, "")
		}
	}
	return record
}

func toTransposedRecord(results []KeyValue, key string, header string) []string {
	record := make([]string, 0, len(results)+1)
	record = append(record, header)
	for _, result := range results {
		if value, ok := result[key]; ok {
			record = append(record, toString(value))
		} else {
			record = append(record, "")
		}
	}
	return record
}
