package notification

import (
	"context"
)

type Service struct {
	repo *Repository
}

/**
 */
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

/**
 */
func (s *Service) GetAll(ctx context.Context) ([]Notification, error) {
	return s.repo.GetAll(ctx)
}

/**
 */
func (s *Service) SaveNotification(ctx context.Context, notification Notification) (Notification, error) {
	return Notification{}, nil
}
