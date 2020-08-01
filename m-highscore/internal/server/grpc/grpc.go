package grpc

import (
	"context"

	pbhighscore "github.com/Vishal-Gupta19/go_grpc_project/m-apis/version1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Grpc sturcture
type Grpc struct {
	address string
	srv     *grpc.Server
}

// HighScore : Take any high value to start
var HighScore = 9999999999.0

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
	log.Info().Msg("SetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}
