package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/auth-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/auth-svc/pkg/pb"
	"github.com/heritechie/go-book-catalog/internal/util"
)

func (server *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	err := util.ValidPasswordConfirmation(req.Password, req.PasswordConfirmation)

	if err != nil {
		return nil, err
	}

	newUser := &models.User{}

	result := server.gorm.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	newAuthUser := &models.UserAuth{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		UserID:         newUser.ID,
	}

	result = server.gorm.Create(&newAuthUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.SignUpResponse{
		Username: newAuthUser.Username,
		UserId:   newAuthUser.UserID.String(),
	}, nil
}
