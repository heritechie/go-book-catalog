package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var book models.Book
	result := server.gorm.Unscoped().First(&book, "id = ?", req.BookId)

	if result.Error != nil {
		return nil, result.Error
	}

	book.Title = req.Title
	book.Description = req.Description
	book.Year = req.Year
	book.PublisherID = uuid.Must(uuid.Parse(req.PublisherId))
	book.AuthorID = uuid.Must(uuid.Parse(req.AuthorId))
	server.gorm.Save(&book)

	return &pb.UpdateBookResponse{
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
