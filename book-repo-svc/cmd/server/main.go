package main

import (
	"net"
	"os"

	"github.com/heritechie/go-book-catalog/book-repo-svc/api"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"

	sharedUtil "github.com/heritechie/go-book-catalog/util"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	sharedConfig, err := sharedUtil.LoadConfig("../")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load shared config")
	}

	if sharedConfig.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load service config")
	}
	runGrpcServer(sharedConfig)

}

func runGrpcServer(config sharedUtil.SharedConfig) {
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(api.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterBookCatalogServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.BookRepoAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
