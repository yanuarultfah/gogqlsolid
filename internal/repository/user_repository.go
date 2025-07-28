package repository

import (
	"context"
	"database/sql"
	"gogqlgensolid/internal/db"
	"gogqlgensolid/internal/model"
)

type UserRepository interface {
	GetByID(id int, ctx context.Context) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(id int, ctx context.Context) error
	ListUsers(offset, limit int, ctx context.Context) ([]*model.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{
		db: db.DB,
	}
}

func (r *userRepo) GetByID(id int, ctx context.Context) (*model.User, error) {
	var user model.User
	query := "select id, name, email from usersgql where id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, err // Other error
	}
	return &user, nil // User found
}
func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "insert into usersgql (id,name, email) values ($1, $2, $3)", user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "update usersgql set name = $1, email = $2 where id = $3", user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepo) Delete(id int, ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "delete from usersgql where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) ListUsers(offset, limit int, ctx context.Context) ([]*model.User, error) {
	query := "select id, name, email from usersgql order by id offset $1 limit $2"
	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
