package server

import (
	"context"

	"github.com/midaef/emmet-server/internal/api"
)

func (s *GRPCServer) AuthWithCredentials(ctx context.Context, req *api.AuthWithCredentialsRequest) (*api.AuthResponseAccessToken, error)  {
	return &api.AuthResponseAccessToken{}, nil
}