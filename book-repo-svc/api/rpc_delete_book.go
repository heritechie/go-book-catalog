package api

import (
	"context"
	"fmt"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
)

func (server *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	var book models.Book
	result := server.gorm.Unscoped().First(&book, "id = ?", req.BookId)

	if result.Error != nil {
		return nil, result.Error
	}

	server.gorm.Delete(&book)

	msg := fmt.Sprintf("Book titled '%s' is deleted", book.Title)

	return &pb.DeleteBookResponse{
		Message: msg,
	}, nil
}
