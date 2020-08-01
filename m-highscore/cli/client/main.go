package main

import (
	"flag"

	pbhighscore "github.com/Vishal-Gupta19/go_grpc_project/m-apis/version1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "locahost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to dial m-highscore gRPC service")
	}

	c := pbhighscore.NewGameClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to get a response")
	}

	if r != nil {
		log.Info().Interface("hinghscore", r.GetHighScore()).Msg("Highscore from m-highscore microservice")
	} else {
		log.Error().Msg("Couldnot get highscore")
	}
}
