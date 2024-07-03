package gevm

import (
	"github.com/ethereum/go-ethereum/common"
)

type TransientStorage struct {
	data map[int]common.Hash
}

func (s *TransientStorage) Load(key int) common.Hash {
	if _, ok := s.data[key]; !ok {
		return common.HexToHash("0x00")
	}
	return s.data[key]
}

func (s *TransientStorage) Store(key int, value common.Hash) (isWarm bool) {
	prev := s.data[key]
	s.data[key] = value
	isWarm = prev == value
	return isWarm
}

func (s *TransientStorage) Clear() {
	//lint:ignore SA4006 ignore unused code warning for this variable
	s = NewTransientStorage()
}

func NewTransientStorage() *TransientStorage {
	return &TransientStorage{
		data: make(map[int]common.Hash),
	}
}
