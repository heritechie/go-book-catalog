package api

import (
	"context"

	"github.com/heritechie/go-book-catalog/auth-svc/internal/db/models"
	"github.com/heritechie/go-book-catalog/auth-svc/pkg/pb"
	"github.com/heritechie/go-book-catalog/internal/util"
)

func (server *Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	var userAuth models.UserAuth
	result := server.gorm.First(&userAuth, "username = ?", req.Username)

	if result.Error != nil {
		return nil, result.Error
	}

	err := util.CheckPassword(req.Password, userAuth.HashedPassword)

	if err != nil {
		return nil, err
	}

	accessToken, err := server.tokenMaker.CreateToken(
		userAuth.Username,
		server.config.AccessTokenDuration,
	)

	return &pb.SignInResponse{
		Username:    userAuth.Username,
		AccessToken: accessToken,
		User: &pb.User{
			Id: userAuth.UserID.String(),
		},
	}, nil
}
