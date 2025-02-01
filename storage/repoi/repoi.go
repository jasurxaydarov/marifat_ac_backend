package repoi

import (
	"context"

	"github.com/jasurxaydarov/marifat_ac_backend/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.UserReq) (*models.User, error)
	GetUser(ctx context.Context, req string) (*models.User, error)
	GetUsers(ctx context.Context, req models.GetList) (*models.UserList, error)
	IsExists(ctx context.Context,req models.IsExists)(*models.IsExistsResp,error)
	LogIn(ctx context.Context, req models.UserLogIn) (*models.Claim, error) 
}

type TeacherRepoI interface {
	CreateUser(ctx context.Context, req models.UserReq) (*models.User, error)
	GetUser(ctx context.Context, req models.Id) (*models.User, error)
}
