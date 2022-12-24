package mockskvtorage

import (
	"context"
	"errors"
	"strings"
)

type Storage struct {
	m map[string][]byte
}

func New() *Storage {
	return &Storage{
		m: make(map[string][]byte),
	}
}

func (s *Storage) GetKeysByPattern(
	ctx context.Context,
	pattern string,
) (res []string, err error) {
	for k := range s.m {
		if strings.Contains(k, pattern) {
			res = append(res, k)
		}
	}

	return
}

func (s *Storage) GetValueByKey(
	ctx context.Context,
	key string,
) (res []byte, err error) {
	var ok bool
	if res, ok = s.m[key]; ok {
		return
	}
	err = errors.New("not found")
	return
}

func (s *Storage) SetValueByKey(
	ctx context.Context,
	key string, value []byte) (err error) {
	s.m[key] = value
	return
}

func (s *Storage) DeleteByKey(
	ctx context.Context,
	key string,
) (err error) {
	var ok bool
	if _, ok = s.m[key]; ok {
		return
	}
	delete(s.m, key)
	err = errors.New("not found")
	return
}
