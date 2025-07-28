package usecase

import (
	"context"
	"gogqlgensolid/internal/model"
	"gogqlgensolid/internal/repository"
)

type UserUseCase interface {
	GetUserByID(id int, ctx context.Context) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(id int, ctx context.Context) error
	ListUsers(offset, limit int, ctx context.Context) ([]*model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: r,
	}
}

func (u *userUseCase) GetUserByID(id int, ctx context.Context) (*model.User, error) {
	return u.repo.GetByID(id, ctx)
}
func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) error {
	return u.repo.Create(ctx, user)
}
func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	return u.repo.Update(ctx, user)
}
func (u *userUseCase) DeleteUser(id int, ctx context.Context) error {
	return u.repo.Delete(id, ctx)
}
func (u *userUseCase) ListUsers(offset, limit int, ctx context.Context) ([]*model.User, error) {
	return u.repo.ListUsers(offset, limit, ctx)
}
