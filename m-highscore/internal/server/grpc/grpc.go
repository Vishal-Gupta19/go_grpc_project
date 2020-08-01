package grpc

import (
	"context"
	"net"

	pbhighscore "github.com/Vishal-Gupta19/go_grpc_project/m-apis/version1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Grpc sturcture
type Grpc struct {
	address string
	srv     *grpc.Server
}

// HighScore : Take any high value to start
var HighScore = 999999999.0

// NewServer ...
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// SetHighScore API
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

// GetHighScore API
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
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
	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("Starting gRPC server for m-highscore microservice")

	// Start serving
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice")
	}

	return nil
}
