package main

import "github.com/Jimbo8702/toll-calc/types"

type Storer interface {
	Insert(types.Distance) error 
}

type MemoryStore struct {}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Insert(d types.Distance) error {
	return nil
}