package main

import (
	"flag"

	"time"

	pbgameengine "github.com/Vishal-Gupta19/go_grpc_project/m-apis/m-game-engine/version1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to dial m-game-engine gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	c := pbgameengine.NewGameEngineClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetSize(timeoutCtx, &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to get a response")
	} else if r != nil {
		log.Info().Interface("highscore", r.GetSize()).Msg("Size from m-game-engine microservice")
	} else {
		log.Error().Msg("Couldnot get size")
	}
}
