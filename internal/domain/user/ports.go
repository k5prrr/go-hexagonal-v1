package user

import (
	"context"
)

type UserRepository interface {
	All(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) (string, error)
	Read(ctx context.Context, uid string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, uid string) error
}
