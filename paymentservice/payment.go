package paymentservice

import "context"

type PaymentService interface {
	Refund(ctx context.Context, request RequestRefund) (response ResponseRefund, code int, err error)
}
