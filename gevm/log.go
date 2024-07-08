package gevm

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

//lint:ignore U1000 ignore unused fields
type Log struct {
	topics []common.Hash
	data   []byte
}

type LogRecord []Log

func (l *LogRecord) AddLog(topics []common.Hash, data []byte) {
	*l = append(*l, Log{topics: topics, data: data})
}

func (l *LogRecord) String() string {
	var sb strings.Builder
	for i, log := range *l {
		sb.WriteString(fmt.Sprintf("Log %d:\n", i))
		sb.WriteString("  Topics:\n")
		for j, topic := range log.topics {
			sb.WriteString(fmt.Sprintf("    Topic %d: %s\n", j, topic.Hex()))
		}
		sb.WriteString(fmt.Sprintf("  Data: %x\n", log.data))
	}
	return sb.String()
}

func NewLogRecord() *LogRecord {
	return new(LogRecord)
}
