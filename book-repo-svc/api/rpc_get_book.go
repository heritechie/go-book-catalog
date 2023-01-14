package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetBook(_ context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	var book models.Book
	result := server.gorm.Unscoped().Find(&book, "id = ?", req.BookId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.GetBookResponse{
		Book: &pb.Book{
			Id:          book.ID.String(),
			Title:       book.Title,
			Description: book.Description,
			Year:        book.Year,
			Author: &pb.Author{
				Id: book.AuthorID.String(),
			},
			Publisher: &pb.Publisher{
				Id: book.PublisherID.String(),
			},
			CreatedAt: timestamppb.New(book.CreatedAt),
			UpdatedAt: timestamppb.New(book.UpdatedAt),
		},
	}, nil
}
