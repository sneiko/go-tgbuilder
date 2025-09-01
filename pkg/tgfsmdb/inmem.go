package tgfsmdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/sneiko/go-tgbuilder/pkg/tgfsm"
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
func (r *InMem) GetUser(_ context.Context, id string) (*tgfsm.User, error) {
	d, ok := r.users.Load(id)
	if ok {
		return d.(*tgfsm.User), nil
	}

	return nil, fmt.Errorf("user not found")
}

// SaveKeyFrame to db
func (r *InMem) SaveKeyFrame(_ context.Context, id, keyFrame string) error {
	r.users.Store(id, &tgfsm.User{ID: id, KeyFrame: keyFrame})
	return nil
}
