package api

import (
	"context"
	"fmt"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
)

func (server *Server) DeletePublisher(ctx context.Context, req *pb.DeletePublisherRequest) (*pb.DeletePublisherResponse, error) {
	var publisher models.Publisher
	result := server.gorm.First(&publisher, "id = ?", req.PublisherId)

	if result.Error != nil {
		return nil, result.Error
	}

	server.gorm.Delete(&publisher)

	msg := fmt.Sprintf("Publisher '%s' is deleted", publisher.Name)

	return &pb.DeletePublisherResponse{
		Message: msg,
	}, nil
}
