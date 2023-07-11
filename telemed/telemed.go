package telemed

import "context"

// Telemed abstraction
type Telemed interface {
	RegisterUserSendBird(ctx context.Context, request *RequestRegisterUserSendBird) (response ResponseRegisterUserSendBird, code int, err error)
}
