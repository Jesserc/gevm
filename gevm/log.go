package gevm

import "github.com/ethereum/go-ethereum/common"

//lint:ignore U1000 ignore unused fields
type Log struct {
	topics []common.Hash
	data   []byte
}

type LogRecord []Log

func (l *LogRecord) AddLog(topics []common.Hash, data []byte) {
	*l = append(*l, Log{topics: topics, data: data})
}

func NewLogRecord() *LogRecord {
	return new(LogRecord)
}
