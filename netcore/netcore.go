package netcore

import "context"

// Netcore abstraction
type Netcore interface {
	AddActivity(ctx context.Context, request []AddActivityRequest) (response AddActivityResponse, err error)
}
