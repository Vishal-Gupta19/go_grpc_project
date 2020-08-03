package grpc

import (
	"context"
	"net"

	pbgameengine "github.com/Vishal-Gupta19/go_grpc_project/m-apis/m-game-engine/version1"
	"github.com/Vishal-Gupta19/go_grpc_project/m-game-engine/internal/server/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Grpc structure for game-engine
type Grpc struct {
	address string
	srv     *grpc.Server
}

// NewServer : Create server for game-engine
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// GetSize ...
func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

// SetScore ...
func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}

// ListenAndServe : To start the service
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to open tcp port")
	}

	// Initialize the server
	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	// Register Grpc struct with Server
	pbgameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("Starting gRPC server for m-game-engine microservice")

	// Start serving
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-game-engine microservice")
	}
	return nil
}
