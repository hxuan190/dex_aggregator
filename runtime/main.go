package main

import (
	"github.com/alphabatem/common/context"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	ctx, err := context.NewCtx()
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	err = ctx.Run()
	if err != nil {
		log.Fatal().Err(err)
		return
	}
}
