package userservice

import (
	"context"
)

// UserService abstraction
type UserService interface {
	Verify(ctx context.Context, request *VerifyRequest) (response VerifyResponse, err error)
	GetAllUser(ctx context.Context, request *FilterUser) (response []UserResponse, err error)
	GetDetailUser(ctx context.Context, request *FilterUser) (response UserResponse, err error)
}
