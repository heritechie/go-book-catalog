package api

import (
	"fmt"

	"github.com/heritechie/go-book-catalog/auth-svc/pkg/pb"
	"github.com/heritechie/go-book-catalog/internal/db"
	sharedUtil "github.com/heritechie/go-book-catalog/internal/util"
	"github.com/heritechie/go-book-catalog/pkg/token"
	"gorm.io/gorm"
)

// Server serves gRPC requests for book repo service.
type Server struct {
	pb.UnimplementedAuthServer
	config     sharedUtil.SharedConfig
	gorm       *gorm.DB
	tokenMaker token.Maker
}

// Init creates a new gRPC server for auth service.
func NewServer(config sharedUtil.SharedConfig) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	conn, err := db.ConnectDB(config)

	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		gorm:       conn,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
