package mqservice

import "context"

// MqService abstraction
type MqService interface {
	SendEmail(ctx context.Context, request *SendEmail) error
	PushNotification(ctx context.Context, request *PushNotification) error
}
