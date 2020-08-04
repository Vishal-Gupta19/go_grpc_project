package bff

import (
	"context"
	"strconv"

	pbgameengine "github.com/Vishal-Gupta19/go_grpc_project/m-apis/m-game-engine/version1"
	pbhighscore "github.com/Vishal-Gupta19/go_grpc_project/m-apis/m-highscore/version1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type gameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameEngineClient
}

// NewGameResource ...
func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

// NewGrpcGameServiceClient to connect with m-highscore
func NewGrpcGameServiceClient(serverAddr string) (pbhighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	}

	log.Info().Msgf("Successfully connected to [%s]", serverAddr)

	if conn == nil {
		log.Info().Msg("m-highscore connection isnil in m-bff")
	}

	client := pbhighscore.NewGameClient(conn)

	return client, nil
}

// NewGrpcGameEngineServiceClient to connect with m-highscore
func NewGrpcGameEngineServiceClient(serverAddr string) (pbgameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	}

	log.Info().Msgf("Successfully connected to [%s]", serverAddr)

	if conn == nil {
		log.Info().Msg("m-game-engine connection isnil in m-bff")
	}

	client := pbgameengine.NewGameEngineClient(conn)

	return client, nil
}

func (gr *gameResource) SetHighScore(c *gin.Context) {
	highscoreString := c.Param("hs") // Key for highscore is hs
	highscoreFloat64, err := strconv.ParseFloat(highscoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert highscore to float")
	}
	gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
		HighScore: highscoreFloat64,
	})
}

func (gr *gameResource) GetHighScore(c *gin.Context) {
	highscoreResponse, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Error while getting highscore")
		return
	}
	hsString := strconv.FormatFloat(highscoreResponse.HighScore, 'e', -1, 64)

	c.JSONP(200, gin.H{
		"hs": hsString,
	})
}

func (gr *gameResource) GetSize(c *gin.Context) {
	sizeResponse, err := gr.gameEngineClient.GetSize(context.Background(), &pbgameengine.GetSizeRequest{})

	if err != nil {
		log.Error().Err(err).Msg("Error while getting size")
	}

	c.JSON(200, gin.H{
		"size": sizeResponse.GetSize(),
	})
}

func (gr *gameResource) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, _ := strconv.ParseFloat(scoreString, 64)

	_, err := gr.gameEngineClient.SetScore(context.Background(), &pbgameengine.SetScoreRequest{
		Score: scoreFloat64,
	})

	if err != nil {
		log.Error().Err(err).Msg("Error while setting score in m-game-engine")
	}
}
