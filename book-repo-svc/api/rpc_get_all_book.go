package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetAllBook(context.Context, *pb.GetAllBookRequest) (*pb.GetAllBookResponse, error) {
	var bookList []models.Book
	result := server.gorm.Unscoped().Find(&bookList)

	if result.Error != nil {
		return nil, result.Error
	}

	var books []*pb.Book

	for _, book := range bookList {
		books = append(books, &pb.Book{
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
		})
	}

	return &pb.GetAllBookResponse{
		Books: books,
	}, nil
}
