package grpc

import (
	"context"
	"google.golang.org/grpc"
	"tech-challenge/pkg/proto/cipher"
	"time"
)

type Client struct {
	grpcServerAddr string
}

func NewClient(grpcServerAddr string) *Client {
	return &Client{
		grpcServerAddr: grpcServerAddr,
	}
}

func (c Client) Encrypt(ctx context.Context, text string, shifts int) (string, error) {
	// @todo make options configurable
	conn, err := grpc.Dial(c.grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := cipher.NewCipherClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := client.Encrypt(ctx, &cipher.EncryptRequest{
		Input: text,
		Shift: int32(shifts),
	})
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}

func (c Client) Decrypt(ctx context.Context, text string, shifts int) (string, error) {
	// @todo avoid copy
	conn, err := grpc.Dial(c.grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := cipher.NewCipherClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := client.Decrypt(ctx, &cipher.DecryptRequest{
		Input: text,
		Shift: int32(shifts),
	})
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}
