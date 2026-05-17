package memory

import (
	"context"
	"time"
)

type Memory interface {
	Load(ctx context.Context, userID string) ([]Entry, error)
	Save(ctx context.Context, userID string, entry Entry) error
	Search(ctx context.Context, userID, query string, k int) ([]Entry, error)
}

type Entry struct {
	Type      EntryType
	Content   string
	CreatedAt time.Time
}

type EntryType string

const (
	EntryPreference EntryType = "preference"
	EntryFact       EntryType = "fact"
	EntryNote       EntryType = "note"
)
