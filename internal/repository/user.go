package repository

import (
	"context"
	"database/sql"
	"gofiber-ewallet/domain"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", db),
	}
}

func (u userRepository) FindById(ctx context.Context, id int64) (user domain.User, err error) {
	data := u.db.From("users").Where(goqu.Ex{"id": id})
	_, err = data.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	data := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})
	_, err = data.ScanStructContext(ctx, &user)
	return
}
