package main

import (
	"flag"
	"github.com/ozoncp/ocp-project-api/internal/tracer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	tracer.InitTracing("ocp_repo_api")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if err := runGrpcAndGateway(); err != nil {
		log.Fatal().Msgf("Something went wrong: %v", err)
	}
}
