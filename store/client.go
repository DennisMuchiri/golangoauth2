package store

import (
	"context"
	"errors"
	"sync"

	"github.com/DennisMuchiri/ke-soundstream-oauth2"
)

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{
		data: make(map[string]oauth2.ClientInfo),
	}
}

// ClientStore client information store
type ClientStore struct {
	sync.RWMutex
	data map[string]oauth2.ClientInfo
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	if c, ok := cs.data[id]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}

// get all current clients information
func (cs *ClientStore) GetAll(ctx context.Context) (map[string]oauth2.ClientInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	c := cs.data

	return c, nil
}

// get all current clients information
func (cs *ClientStore) ReplaceAll(clients map[string]oauth2.ClientInfo, ctx context.Context) (bool, error) {
	cs.RLock()
	defer cs.RUnlock()

	cs.data = clients
	return true, nil
}

// Set set client information
func (cs *ClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()

	cs.data[id] = cli
	return
}
