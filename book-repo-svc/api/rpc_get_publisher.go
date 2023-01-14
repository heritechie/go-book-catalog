package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetPublisher(ctx context.Context, req *pb.GetPublisherRequest) (*pb.GetPublisherResponse, error) {
	var publisher models.Publisher
	result := server.gorm.First(&publisher, "id = ?", req.PublisherId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.GetPublisherResponse{
		Publisher: &pb.Publisher{
			Id:        publisher.ID.String(),
			Name:      publisher.Name,
			CreatedAt: timestamppb.New(publisher.CreatedAt),
			UpdatedAt: timestamppb.New(publisher.UpdatedAt),
		},
	}, nil
}
