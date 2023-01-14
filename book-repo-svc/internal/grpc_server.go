package grpc_server

import (
	"github.com/heritechie/go-book-catalog/util"

	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
)

// Server serves gRPC requests for book repo service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.SharedConfig
}

// Init creates a new gRPC server.
func Init(config util.SharedConfig) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}
