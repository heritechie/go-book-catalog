package api

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	now := time.Now()

	newBook := &models.Book{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		PublisherID: uuid.Must(uuid.Parse(req.GetPublisherId())),
		AuthorID:    uuid.Must(uuid.Parse(req.GetAuthorId())),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := server.gorm.Unscoped().Create(&newBook)

	if result.Error != nil {
		return nil, result.Error
	}

	var book models.Book
	result = server.gorm.Unscoped().First(&book, "id = ?", newBook.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.CreateBookResponse{
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
