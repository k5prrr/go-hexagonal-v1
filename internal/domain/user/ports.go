package user

import (
	"context"
)

type AnyUserService interface {
	UserByUid(uid string) (*User, error)
	CreateUser(user *User) (string, error)
	UpdateUser(user *User) error
	DeleteByUid(uid string) error
	AllUsers() ([]*User, error)
}

type UserRepository interface {
	All(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) (string, error)
	Read(ctx context.Context, uid string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, uid string) error
}
