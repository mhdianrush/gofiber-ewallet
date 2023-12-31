package domain

import (
	"context"
	"gofiber-ewallet/dto"
)

type User struct {
	Id       int64  `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindById(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
}
