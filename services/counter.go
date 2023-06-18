package services

import (
	"github.com/dyuri/templ-counter/db"
	"golang.org/x/exp/slog"
)

// Counter counts the clicks
type Counter struct {
	Logger     *slog.Logger
	CountStore *db.CountStore
}

// NewCounter creates a new counter instance
func NewCounter(logger *slog.Logger, countStore *db.CountStore) *Counter {
	return &Counter{
		Logger:     logger,
		CountStore: countStore,
	}
}
