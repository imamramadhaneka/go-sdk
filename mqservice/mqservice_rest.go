package mqservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

const (
	// TextNewMqService const
	TextNewMqService = "NewMqService"
)

type mqServiceRESTImpl struct {
	host string
}

// NewMqServiceREST constructor
func NewMqServiceREST(host string) MqService {
	return &mqServiceRESTImpl{
		host: host,
	}
}

// SendEmail
func (m *mqServiceRESTImpl) SendEmail(ctx context.Context, request *SendEmail) (err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "MqService:SendEmail")
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
		uri      string
		payload  []byte
		response RestAPIResult
	)

	headers := map[string]string{
		echo.HeaderContentType: echo.MIMEApplicationJSON,
	}

	buff := &bytes.Buffer{}
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(request)
	payload = buff.Bytes()

	uri = fmt.Sprintf("%s/api/v1/emails/blank", m.host)
	body, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		tracer.Log(ctx, "response_not_success", err)
		return err
	}

	json.Unmarshal(body, &response)
	if response.Status != http.StatusOK {
		err = errors.New(response.Errors)
		tracer.Log(ctx, "response_not_success", err)
		return err
	}

	return
}

func (m *mqServiceRESTImpl) PushNotification(ctx context.Context, request *PushNotification) (err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "MqService:PushNotification")
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
		uri      string
		payload  []byte
		response NotificationRestAPIResult
	)

	headers := map[string]string{
		echo.HeaderContentType: echo.MIMEApplicationJSON,
	}

	buff := &bytes.Buffer{}
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(request)
	payload = buff.Bytes()

	uri = fmt.Sprintf("%s/api/v1/push_notifications", m.host)
	body, resCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		tracer.Log(ctx, "response_not_success", err)
		return err
	}

	json.Unmarshal(body, &response)
	if resCode != http.StatusOK {
		err = errors.New(response.Errors)
		tracer.Log(ctx, "response_not_success", err)
		return err
	}

	return
}
