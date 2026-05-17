package transport

import "context"

type Transport interface {
	Name() string
	Start(ctx context.Context, chat ChatFunc) error
	Stop() error
}

type ChatFunc func(ctx context.Context, userID, msg string) (string, error)
