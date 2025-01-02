package fsmdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"tg_star_miner/pkg/tgfsm/fsmmodel"
)

// InMem db
type InMem struct {
	users sync.Map
}

// NewInMem create new db
func NewInMem() *InMem {
	return &InMem{
		users: sync.Map{},
	}
}

// GetUser from db
func (r *InMem) GetUser(_ context.Context, id uuid.UUID) (*fsmmodel.User, error) {
	d, ok := r.users.Load(id)
	if ok {
		return d.(*fsmmodel.User), nil
	}

	return nil, fmt.Errorf("user not found")
}

// SaveKeyFrame to db
func (r *InMem) SaveKeyFrame(_ context.Context, id uuid.UUID, keyFrame string) error {
	r.users.Store(id, &fsmmodel.User{ID: id, KeyFrame: keyFrame})
	return nil
}
