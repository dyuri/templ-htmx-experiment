package db

// CountStore is a simple in-memory counter
type CountStore struct { // TODO use actual db
	Count int
}

// NewCountStore creates a new CountStore instance
func NewCountStore() (s *CountStore, err error) {
	s = &CountStore{}
	return
}
