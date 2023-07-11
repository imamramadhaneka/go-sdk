package paymentservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

type paymentServiceRESTImpl struct {
	host string
	auth string
}

// NewActivityServiceREST constructor
func NewPaymentServiceREST(host string, auth string) PaymentService {
	return &paymentServiceRESTImpl{
		host: host,
		auth: auth,
	}
}

// InsertLog
func (t *paymentServiceRESTImpl) Refund(ctx context.Context, request RequestRefund) (response ResponseRefund, code int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "PaymentService:Refund")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
		if err != nil {
			span.SetError(err)
		}
		span.Finish()
	}()

	var (
		uri     string
		payload []byte
		resp    RestAPIResult
	)

	resp.Data = new(ResponseRefund)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: t.auth,
	}

	payload, _ = json.Marshal(request)

	uri = fmt.Sprintf("%s/cf/refund", t.host)
	body, statusCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		return response, statusCode, err
	}

	json.Unmarshal(body, &resp)
	if results, ok := resp.Data.(*ResponseRefund); ok {
		return *results, statusCode, nil
	}

	return response, statusCode, nil
}
