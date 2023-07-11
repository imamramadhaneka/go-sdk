package userservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

type userServiceRESTImpl struct {
	host string
	auth string
}

// NewUserServiceREST constructor
func NewUserServiceREST(host string, auth string) UserService {
	return &userServiceRESTImpl{
		host: host,
		auth: auth,
	}
}

// InsertLog
func (t *userServiceRESTImpl) Verify(ctx context.Context, request *VerifyRequest) (response VerifyResponse, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "UserService:Verify")

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

	resp.Data = new(VerifyResponse)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: t.auth,
	}

	payload, _ = json.Marshal(request)

	uri = fmt.Sprintf("%s/user/verify", t.host)
	body, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodPost, uri, payload, headers)
	if err != nil {
		return response, err
	}

	json.Unmarshal(body, &resp)
	if results, ok := resp.Data.(*VerifyResponse); ok {
		return *results, nil
	}

	return response, nil

}

// GetAllUser
func (t *userServiceRESTImpl) GetAllUser(ctx context.Context, filter *FilterUser) (response []UserResponse, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "UserService:GetLog")

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
	tags["filter"] = filter

	resp.Data = new([]UserResponse)

	candiTokenClaim := ctx.Value(candishared.ContextKeyTokenClaim).(*candishared.TokenClaim)
	claim := candiTokenClaim.Additional.(VerifyResponse)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: fmt.Sprintf("Bearer %s", claim.Token),
	}

	if filter.Page >= 1 {
		page = filter.Page
	}
	if filter.PerPage >= 1 {
		limit = filter.PerPage
	}
	if strings.TrimSpace(filter.Email) != "" {
		queryPrams += "&email=" + url.QueryEscape(filter.Email)
	}
	if strings.TrimSpace(filter.Status) != "" {
		queryPrams += "&status=" + url.QueryEscape(filter.Status)
	}
	if strings.TrimSpace(filter.Search) != "" {
		queryPrams += "&search=" + url.QueryEscape(filter.Search)
	}

	uri = fmt.Sprintf("%s/user?page=%d&perPage=%d&orderBy=%s&sort=%s%s", t.host, page, limit, filter.OrderBy, filter.Sort, queryPrams)
	b, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return response, err
	}

	json.Unmarshal(b, &resp)

	if userList, ok := reflect.ValueOf(resp.Data).Elem().Interface().([]UserResponse); ok {
		return userList, nil
	}

	return nil, errors.New("invalid user response")
}

// GetDetailUser
func (t *userServiceRESTImpl) GetDetailUser(ctx context.Context, filter *FilterUser) (response UserResponse, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "UserService:GetLogByID")

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
	tags["filter"] = filter

	resp.Data = new(UserResponse)

	candiTokenClaim := ctx.Value(candishared.ContextKeyTokenClaim).(*candishared.TokenClaim)
	claim := candiTokenClaim.Additional.(VerifyResponse)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: fmt.Sprintf("Bearer %s", claim.Token),
	}

	if filter == nil || strings.TrimSpace(filter.UUID) == "" {
		return response, errors.New("parameter uuid is required")
	}

	uri = fmt.Sprintf("%s/user/%s", t.host, filter.UUID)
	b, _, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return response, err
	}

	json.Unmarshal(b, &resp)
	if user, ok := resp.Data.(*UserResponse); ok {
		return *user, nil
	}

	return response, errors.New("invalid user detail response")

}
