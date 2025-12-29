package db

import (
	"github.com/dgraph-io/badger/v3"
)

type Store struct {
	db *badger.DB
}

func NewStore(path string) (*Store, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() { s.db.Close() }

func (s *Store) Set(key string, val []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), val)
	})
}

func (s *Store) Get(key string) ([]byte, error) {
	var val []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, e := txn.Get([]byte(key))
		if e != nil {
			return e
		}
		val, _ = item.ValueCopy(nil)
		return nil
	})
	return val, err
}

func (s *Store) Delete(key string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
