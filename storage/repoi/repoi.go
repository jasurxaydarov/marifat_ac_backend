package repoi

import (
	"context"

	"github.com/jasurxaydarov/marifat_ac_backend/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.UserReq) (*models.User, error)
	GetUser(ctx context.Context, req string) (*models.User, error)
}
