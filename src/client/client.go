package client

import "context"

type Client interface {
	Encrypt(ctx context.Context, text string, shifts int) (string, error)
	Decrypt(ctx context.Context, text string, shifts int) (string, error)
}
