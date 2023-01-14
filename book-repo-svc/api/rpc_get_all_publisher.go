package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetAllPublisher(context.Context, *pb.GetAllPublisherRequest) (*pb.GetAllPublisherResponse, error) {
	var publisherList []models.Publisher
	result := server.gorm.Find(&publisherList)

	if result.Error != nil {
		return nil, result.Error
	}

	var publishers []*pb.Publisher

	for _, publisher := range publisherList {
		publishers = append(publishers, &pb.Publisher{
			Id:        publisher.ID.String(),
			Name:      publisher.Name,
			CreatedAt: timestamppb.New(publisher.CreatedAt),
			UpdatedAt: timestamppb.New(publisher.UpdatedAt),
		})
	}

	return &pb.GetAllPublisherResponse{
		Publishers: publishers,
	}, nil
}
