package tgfsm

import (
	"context"
)

// Storage interface for implementation storage
type Storage interface {
	GetUser(ctx context.Context, id string) (*User, error)
	SaveKeyFrame(ctx context.Context, id string, keyFrame string) error
}

// Manager state machine for telegram
type Manager struct {
	storage Storage
}

// New Manager
func New(storage Storage) *Manager {
	return &Manager{
		storage: storage,
	}
}

// GetUser from db
func (s *Manager) GetUser(ctx context.Context, id string) (*User, error) {
	return s.storage.GetUser(ctx, id)
}

// SaveKeyFrame to db
func (s *Manager) SaveKeyFrame(ctx context.Context, id string, keyFrame string) error {
	return s.storage.SaveKeyFrame(ctx, id, keyFrame)
}
