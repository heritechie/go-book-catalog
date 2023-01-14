package api

import (
	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"github.com/heritechie/go-book-catalog/util"
	"gorm.io/gorm"
)

// Server serves gRPC requests for book repo service.
type Server struct {
	pb.UnimplementedBookCatalogServer
	config util.SharedConfig
	gorm   *gorm.DB
}

// Init creates a new gRPC server.
func NewServer(config util.SharedConfig) (*Server, error) {
	server := &Server{
		config: config,
		gorm:   db.ConnectDB(config),
	}

	return server, nil
}
