package main

import (
	"net"
	"os"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/grpc_server"
	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/util"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"

	sharedUtil "github.com/heritechie/go-book-catalog/util"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	sharedConfig, err := sharedUtil.LoadConfig("../")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load shared config")
	}

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load service config")
	}

	if sharedConfig.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	dsn := "host=localhost user=root password=secret dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	// runDBMigration(config.MigrationURL, sharedConfig.DBSource)
	runGrpcServer(sharedConfig)

}

// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create new migrate instance")
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal().Err(err).Msg("failed to run migrate up")
// 	}

// 	log.Info().Msg("db migrated successfully")
// }

func runGrpcServer(config util.SharedConfig) {
	server, err := grpc_server.Init(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(grpc_server.GrpcLogger)
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
