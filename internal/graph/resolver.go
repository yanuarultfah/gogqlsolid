package graph

import (
	"gogqlgensolid/internal/model"
	"gogqlgensolid/internal/usecase"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	svc usecase.UserUseCase
}

// NewResolver creates a new resolver instance
func NewResolver(userService usecase.UserUseCase) *Resolver {
	return &Resolver{
		svc: userService,
	}
}

func (r *Resolver) ResolverUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, nil // Handle error appropriately
	}
	user, err := r.svc.GetUserByID(id, p.Context)
	if err != nil {
		return nil, err // Handle error appropriately
	}
	return user, nil
}

func (r *Resolver) ResolverUsers(p graphql.ResolveParams) (interface{}, error) {
	offset, okOffset := p.Args["offset"].(int)
	limit, okLimit := p.Args["limit"].(int)
	if !okOffset || !okLimit {
		return nil, nil // Handle error appropriately
	}
	users, err := r.svc.ListUsers(offset, limit, p.Context)
	if err != nil {
		return nil, err // Handle error appropriately
	}
	return users, nil
}

func (r *Resolver) ResolverCreateUser(p graphql.ResolveParams) (interface{}, error) {
	id, okID := p.Args["id"].(int)
	name, okName := p.Args["name"].(string)
	email, okEmail := p.Args["email"].(string)
	if !okID || !okName || !okEmail {
		return nil, nil // Handle error appropriately
	}
	user := &model.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	err := r.svc.CreateUser(p.Context, user)
	if err != nil {
		return nil, err // Handle error appropriately
	}
	return user, nil
}

func (r *Resolver) ResolverUpdateUser(p graphql.ResolveParams) (interface{}, error) {
	id, okID := p.Args["id"].(int)
	name, okName := p.Args["name"].(string)
	email, okEmail := p.Args["email"].(string)
	if !okID || !okName || !okEmail {
		return nil, nil // Handle error appropriately
	}
	user := &model.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	err := r.svc.UpdateUser(p.Context, user)
	if err != nil {
		return nil, err // Handle error appropriately
	}
	return user, nil
}

func (r *Resolver) ResolverDeleteUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, nil // Handle error appropriately
	}
	err := r.svc.DeleteUser(id, p.Context)
	if err != nil {
		return nil, err // Handle error appropriately
	}
	return "User deleted successfully", nil
}
