package storage

import (
	"context"

	"github.com/luqus/s/types"
)

type AuthenticationStorage interface {
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	GetUserByUID(ctx context.Context, uid string) (*types.User, error)
	CheckIfEmailExists(ctx context.Context, email string) (int64, error)
	CheckIfPhoneNumberExists(ctx context.Context, phoneNumber string) (int64, error)
	CreateUser(ctx context.Context, user *types.User) error
}

type AlertStorage interface{}
