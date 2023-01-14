package api

import (
	"context"
	"time"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreatePublisher(ctx context.Context, req *pb.CreatePublisherRequest) (*pb.CreatePublisherResponse, error) {
	now := time.Now()

	newPublisher := &models.Publisher{
		Name:      req.GetName(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := server.gorm.Create(&newPublisher)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.CreatePublisherResponse{
		Publisher: &pb.Publisher{
			Id:        newPublisher.ID.String(),
			Name:      newPublisher.Name,
			CreatedAt: timestamppb.New(newPublisher.CreatedAt),
			UpdatedAt: timestamppb.New(newPublisher.UpdatedAt),
		},
	}, nil
}
