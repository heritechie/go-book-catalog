package api

import (
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"github.com/heritechie/go-book-catalog/internal/db"
	"github.com/heritechie/go-book-catalog/internal/util"
	"gorm.io/gorm"
)

// Server serves gRPC requests for book repo service.
type Server struct {
	pb.UnimplementedBookCatalogServer
	config util.SharedConfig
	gorm   *gorm.DB
}

// Init creates a new gRPC server for book catalog repo service.
func NewServer(config util.SharedConfig) (*Server, error) {

	conn, err := db.ConnectDB(config)

	if err != nil {
		return nil, err
	}

	server := &Server{
		config: config,
		gorm:   conn,
	}

	return server, nil
}
