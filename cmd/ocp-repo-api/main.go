package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozoncp/ocp-project-api/internal/config"
	"github.com/ozoncp/ocp-project-api/internal/tracer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	debug      = flag.Bool("debug", false, "sets log level to debug")
	configPath = flag.String("config", "", "sets config path")
)

func main() {
	flag.Parse()
	config.LoadGlobal(*configPath)

	tracer.InitTracing("ocp_repo_api")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	go func() {
		// See comment in ocp-project-api
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		log.Info().Msgf("Caught signal: %v", sig)
		GrpcMutex.Lock()
		defer GrpcMutex.Unlock()
		if GrpcServer != nil {
			log.Info().Msg("Stopping Grpc server...")
			GrpcServer.GracefulStop()
			log.Info().Msg("Grpc server is stopped")
		} else {
			log.Info().Msg("No Grpc server is started")
		}
		syscall.Exit(1)
	}()

	if err := runGrpcAndGateway(); err != nil {
		log.Fatal().Msgf("Something went wrong: %v", err)
	}
}
