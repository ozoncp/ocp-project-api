package main

import (
	"flag"
	"net/http"
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

	log.Debug().Msg("Debug level is up")

	go func() {
		const pprofEndpoint = ":8090"
		log.Info().Msgf("Profiling service on %[1]s (to watch go to %[1]s/debug/pprof link)", pprofEndpoint)
		if err := http.ListenAndServe(pprofEndpoint, nil); err != nil {
			log.Warn().Msgf("Profiling service failed: %v", err)
		}
	}()

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
