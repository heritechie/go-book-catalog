package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetAuthor(ctx context.Context, req *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	var author models.Author
	result := server.gorm.First(&author, "id = ?", req.AuthorId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.GetAuthorResponse{
		Author: &pb.Author{
			Id:        author.ID.String(),
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CreatedAt),
			UpdatedAt: timestamppb.New(author.UpdatedAt),
		},
	}, nil
}
