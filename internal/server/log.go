package server

import (
	"fmt"
	"sync"
)

var ErrorOffsetNotFound = fmt.Errorf("offset not fount")

type Log struct {
	mu      sync.Mutex
	records []Record
}

type Record struct {
	Offset uint64
	Value  []byte
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	append(c.records, record)
	return record.Offset, nil
}
