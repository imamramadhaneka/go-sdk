package netcore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

type netcoreRESTImpl struct {
	host string
	auth string
}

// NewNetcoreRESTImpl constructor
func NewNetcoreRESTImpl(host string, auth string) Netcore {
	return &netcoreRESTImpl{
		host: host,
		auth: auth,
	}
}

func (t *netcoreRESTImpl) AddActivity(ctx context.Context, request []AddActivityRequest) (response AddActivityResponse, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "Netcore:AddActivity")

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
		resp    AddActivityResponse
	)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: fmt.Sprintf("bearer %s", t.auth),
	}
	
	payload, _ = json.Marshal(request)

	uri = fmt.Sprintf("%s/activity/upload", t.host)
	body, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)

	_  = json.Unmarshal(body, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
