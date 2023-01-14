package api

import (
	"context"
	"fmt"

	"github.com/heritechie/go-book-catalog/book-repo-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/book-repo-svc/pkg/pb"
)

func (server *Server) DeleteAuthor(ctx context.Context, req *pb.DeleteAuthorRequest) (*pb.DeleteAuthorResponse, error) {
	var author models.Author
	result := server.gorm.First(&author, "id = ?", req.AuthorId)

	if result.Error != nil {
		return nil, result.Error
	}

	server.gorm.Delete(&author)

	msg := fmt.Sprintf("Author '%s' is deleted", author.Name)

	return &pb.DeleteAuthorResponse{
		Message: msg,
	}, nil
}
