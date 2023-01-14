package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetAllAuthor(context.Context, *pb.GetAllAuthorRequest) (*pb.GetAllAuthorResponse, error) {
	var authorList []models.Author
	result := server.gorm.Find(&authorList)

	if result.Error != nil {
		return nil, result.Error
	}

	var authors []*pb.Author

	for _, author := range authorList {
		authors = append(authors, &pb.Author{
			Id:        author.ID.String(),
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CreatedAt),
			UpdatedAt: timestamppb.New(author.UpdatedAt),
		})
	}

	return &pb.GetAllAuthorResponse{
		Authors: authors,
	}, nil
}
