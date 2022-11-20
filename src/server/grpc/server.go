package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tech-challenge/pkg/proto/cipher"
	"tech-challenge/server"
)

type Server struct {
	cipher.UnimplementedCipherServer
	cipher server.Cipher
}

func NewServer(cipher server.Cipher) *Server {
	return &Server{cipher: cipher}
}

func (s *Server) Encrypt(ctx context.Context, req *cipher.EncryptRequest) (*cipher.EncryptResponse, error) {
	if req.Shift == 0 {
		return nil, status.Error(codes.InvalidArgument, "shift amount should not be a zero")
	}
	return &cipher.EncryptResponse{
		Result: s.cipher.Encrypt(req.Input, int(req.Shift)),
	}, nil
}
func (s *Server) Decrypt(ctx context.Context, req *cipher.DecryptRequest) (*cipher.DecryptResponse, error) {
	if req.Shift == 0 {
		return nil, status.Error(codes.InvalidArgument, "shift amount should not be a zero")
	}
	return &cipher.DecryptResponse{
		Result: s.cipher.Decrypt(req.Input, int(req.Shift)),
	}, nil
}
