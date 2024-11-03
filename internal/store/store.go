package store

import (
	"fmt"

	"github.com/schollz/jsonstore"
)

type Store struct {
	Path  string
	store *jsonstore.JSONStore
}

func NewStore(path string) *Store {
	return &Store{Path: path}
}

func (s *Store) OpenStore() error {
	var err error
	s.store, err = jsonstore.Open(s.Path)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Persist() error {
	return jsonstore.Save(s.store, s.Path)
}

func (s *Store) StoreRepoPath(path string) {
	key := fmt.Sprintf("repo:%d", len(s.store.Keys()))
	s.store.Set(key, path)
	s.Persist()
}

func (s *Store) GetRepoPaths() []string {
	paths := make([]string, 0)
	for _, key := range s.store.Keys() {
		var path string
		if err := s.store.Get(key, &path); err == nil {
			paths = append(paths, path)
		}
	}
	return paths
}

func (s *Store) RemoveRepoPath(path string) error {
	// Find and remove the key for this path
	for _, key := range s.store.Keys() {
		var storedPath string
		if err := s.store.Get(key, &storedPath); err == nil {
			if storedPath == path {
				s.store.Delete(key)
				return s.Persist()
			}
		}
	}
	return fmt.Errorf("path not found in store")
}
