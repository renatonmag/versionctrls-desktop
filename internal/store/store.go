package store

import "github.com/schollz/jsonstore"

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

func (s *Store) SaveStore() error {
	return jsonstore.Save(s.store, s.Path)
}
