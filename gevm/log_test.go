package gevm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLogRecord(t *testing.T) {
	tests := []struct {
		name       string
		operations func(l *LogRecord) string
		expected   string
	}{
		{
			name: "AddLog and String",
			operations: func(l *LogRecord) string {
				topics := []common.Hash{common.HexToHash("0x1"), common.HexToHash("0x2")}
				data := []byte{0x01, 0x02, 0x03}
				l.AddLog(topics, data)
				return l.String()
			},
			expected: "Log 0:\n  Topics:\n    Topic 0: 0x0000000000000000000000000000000000000000000000000000000000000001\n    Topic 1: 0x0000000000000000000000000000000000000000000000000000000000000002\n  Data: 010203\n",
		},
		{
			name: "Add multiple logs and String",
			operations: func(l *LogRecord) string {
				l.AddLog([]common.Hash{common.HexToHash("0x1")}, []byte{0x01})
				l.AddLog([]common.Hash{common.HexToHash("0x2")}, []byte{0x02})
				return l.String()
			},
			expected: "Log 0:\n  Topics:\n    Topic 0: 0x0000000000000000000000000000000000000000000000000000000000000001\n  Data: 01\nLog 1:\n  Topics:\n    Topic 0: 0x0000000000000000000000000000000000000000000000000000000000000002\n  Data: 02\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logRecord := NewLogRecord()
			result := tt.operations(logRecord)
			assert.Equal(t, tt.expected, result)
		})
	}
}
