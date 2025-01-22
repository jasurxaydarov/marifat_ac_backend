package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
	"github.com/jasurxaydarov/marifat_ac_backend/storage/repoi"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserREpo(db *pgx.Conn) repoi.UserRepoI {

	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(ctx context.Context, req models.UserReq) (*models.User, error) {
	var id uuid.UUID

	query := `
		INSERT INTO users(
			user_id,
			user_name,
			user_email,
			user_password,
			description
		)VALUES(
			$1, $2, $3, $4, $5
		)
	
		`

	id = uuid.New()

	_, err := u.db.Exec(ctx,
		query,
		id,
		req.User_name,
		req.User_email,
		req.Password,
		req.Description,
	)

	if err != nil {
		fmt.Println("err on create user", err)
		return nil, err
	}

	resp, err := u.GetUser(ctx, id.String())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetUser(ctx context.Context, req string) (*models.User, error) {

	var resp models.User
	query := `
		SELECT 
			user_id,
			user_name,
			user_email,
			user_password,
			description
		FROM
			users
		WHERE
			user_id = $1

	`

	row := u.db.QueryRow(ctx, query, req)

	err := row.Scan(
		&resp.User_id,
		&resp.User_name,
		&resp.User_email,
		&resp.Password,
		&resp.Description,
	)

	if err != nil {

		fmt.Println("err on getting user", err)
		return nil, err
	}

	return &resp, nil
}
