package api

import (
	"context"
	"time"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAuthor(ctx context.Context, req *pb.CreateAuthorRequest) (*pb.CreateAuthorResponse, error) {
	now := time.Now()

	newAuthor := &models.Author{
		Name:      req.GetName(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := server.gorm.Create(newAuthor)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.CreateAuthorResponse{
		Author: &pb.Author{
			Name:      newAuthor.Name,
			CreatedAt: timestamppb.New(newAuthor.CreatedAt),
			UpdatedAt: timestamppb.New(newAuthor.UpdatedAt),
		},
	}, nil
}
