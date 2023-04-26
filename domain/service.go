package domain

import (
	"context"
	"fmt"
)

type RecordService interface {
	GetRecords(ctx context.Context, id int) ([]*Record, error)
}

type recordService struct {
	m map[int][]*Record
}

func NewRecordService(rs []*Record) RecordService {
	m := make(map[int][]*Record)

	for _, v := range rs {
		m[v.Id] = append(m[v.Id], v)
	}

	return &recordService{
		m: m,
	}
}

func (s *recordService) GetRecords(ctx context.Context, id int) ([]*Record, error) {
	rs, ok := s.m[id]
	if !ok {
		return nil, fmt.Errorf("no records found with id %v", id)
	}

	return rs, nil
}
