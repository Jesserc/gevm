package gevm

import "slices"

type Storage struct {
	data  map[int][]byte
	cache []int // slot keys cache for warm storage access
}

func (s *Storage) Load(key int) (bool, []byte) {
	warmAccess := slices.Contains(s.cache, key)
	if !warmAccess {
		s.cache = append(s.cache, key)
	}

	if _, ok := s.data[key]; !ok {
		return false, []byte{0x00}
	}
	return warmAccess, s.data[key]
}

func (s *Storage) Store(key int, value []byte) {
	s.data[key] = value
}

func NewStorage() *Storage {
	return &Storage{data: make(map[int][]byte), cache: make([]int, 0)}
}
