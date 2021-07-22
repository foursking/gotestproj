package service

import (
	"context"
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/dao"
)

// Service app's service logic
type Service struct {
	dao dao.Dao
}

// New creates service
func New(dao dao.Dao) *Service {
	srv := Service{
		dao: dao,
	}
	return &srv
}

// Ping ping service
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close service
func (s *Service) Close() error {
	return s.dao.Close()
}