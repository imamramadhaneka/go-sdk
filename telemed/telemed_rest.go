package telemed

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

type telemedRESTImpl struct {
	host string
}

// NewTelemedREST constructor
func NewTelemedREST(host string) Telemed {
	return &telemedRESTImpl{
		host: host,
	}
}

// RegisterUserSendBird
func (t *telemedRESTImpl) RegisterUserSendBird(ctx context.Context, request *RequestRegisterUserSendBird) (response ResponseRegisterUserSendBird, code int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "Telemed:RegisterUserSendBird")

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

	resp.Data = new(ResponseRegisterUserSendBird)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: AuthorizationToken,
	}

	payload, _ = json.Marshal(request)

	uri = fmt.Sprintf("%s/inner-services/sb-register", t.host)
	body, statusCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		return response, statusCode, err
	}

	json.Unmarshal(body, &resp)
	if results, ok := resp.Data.(*ResponseRegisterUserSendBird); ok {
		return *results, statusCode, nil
	}

	return response, statusCode, nil

}
