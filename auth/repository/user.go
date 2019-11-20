package repository

import (
	"app/auth"
	"app/db/model"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Get(ctx context.Context, username, password string) (*auth.User, error) {
	u, err := model.Users(
		qm.Where("username=? and password=?", username, password)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.toUser(ctx, u), nil
}

func (r *UserRepository) Create(ctx context.Context, username, password string) error {
	u := model.User{
		UserID:   uuid.UUID{}.String(),
		Username: username,
		Password: password,
	}

	if err := u.Insert(ctx, r.db, boil.Infer()); err != nil {
		return errors.Wrap(err, "can't create user")
	}
	return nil
}

func (r *UserRepository) toUser(ctx context.Context, u *model.User) *auth.User {
	return &auth.User{
		ID:       u.ID,
		UserID:   uuid.MustParse(u.UserID),
		Username: u.Username,
	}
}
