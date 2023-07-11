package activityservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

type activityServiceRESTImpl struct {
	host string
	auth string
}

// NewActivityServiceREST constructor
func NewActivityServiceREST(host string, auth string) ActivityService {
	return &activityServiceRESTImpl{
		host: host,
		auth: auth,
	}
}

// InsertLog
func (t *activityServiceRESTImpl) InsertLog(ctx context.Context, request *RequestInsertLog) (response ResponseLog, code int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "ActivityService:InsertLog")

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

	resp.Data = new(ResponseLog)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: t.auth,
	}

	payload, _ = json.Marshal(request)

	uri = fmt.Sprintf("%s/logs", t.host)
	body, statusCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		return response, statusCode, err
	}

	json.Unmarshal(body, &resp)
	if results, ok := resp.Data.(*ResponseLog); ok {
		return *results, statusCode, nil
	}

	return response, statusCode, nil

}

// GetLog
func (t *activityServiceRESTImpl) GetLog(ctx context.Context, param *LogParam) (response []ResponseLog, totaldata int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "ActivityService:GetLog")

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
		uri        string
		resp       RestAPIResult
		queryPrams string
		page       = 1
		limit      = 10
	)

	tags := span.Tags()
	tags["param"] = param

	resp.Data = new([]ResponseLog)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: t.auth,
	}

	if param.Page >= 0 {
		page = param.Page
	}
	if param.PerPage >= 0 {
		limit = param.PerPage
	}
	if strings.TrimSpace(param.Module) != "" {
		queryPrams += "&module=" + url.QueryEscape(param.Module)
	}
	if strings.TrimSpace(param.Service) != "" {
		queryPrams += "&service=" + url.QueryEscape(param.Service)
	}
	if strings.TrimSpace(param.Target) != "" {
		queryPrams += "&target=" + url.QueryEscape(param.Target)
	}
	if strings.TrimSpace(param.Action) != "" {
		queryPrams += "&action=" + url.QueryEscape(param.Action)
	}
	if strings.TrimSpace(param.PaymentMethodSlug) != "" {
		queryPrams += "&paymentMethodSlug=" + url.QueryEscape(param.PaymentMethodSlug)
	}

	uri = fmt.Sprintf("%s/logs?page=%d&perPage=%d&orderBy=%s&sort=%s%s", t.host, page, limit, param.OrderBy, param.Sort, queryPrams)
	b, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return response, 0, err
	}

	json.Unmarshal(b, &resp)

	if logList, ok := reflect.ValueOf(resp.Data).Elem().Interface().([]ResponseLog); ok {
		metaResp := resp.Meta.(map[string]interface{})
		metaPaginationResp := metaResp["pagination"].(map[string]interface{})
		metaPaginationTotalResp := int(metaPaginationResp["total"].(float64))
		return logList, metaPaginationTotalResp, nil
	}

	return nil, 0, errors.New("invalid log response")
}

// GetLog
func (t *activityServiceRESTImpl) GetLogByID(ctx context.Context, param *LogParam) (response ResponseLog, totaldata int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "ActivityService:GetLogByID")

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
		uri  string
		resp RestAPIResult
	)

	tags := span.Tags()
	tags["param"] = param

	resp.Data = new(ResponseLog)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: t.auth,
	}

	if param == nil || strings.TrimSpace(param.ID) == "" {
		return response, 0, errors.New("parameter id is required")
	}

	uri = fmt.Sprintf("%s/logs/%s", t.host, param.ID)
	b, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return response, 0, err
	}

	json.Unmarshal(b, &resp)
	if log, ok := resp.Data.(*ResponseLog); ok {
		return *log, 1, nil
	}

	return response, http.StatusInternalServerError, errors.New("invalid log detail response")

}
