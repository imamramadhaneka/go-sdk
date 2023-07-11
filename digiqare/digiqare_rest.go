package digiqare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/labstack/echo"
)

const (
	// TextNewDigiqare const
	TextNewDigiqare = "NewDigiqare"
)

type digiqareRESTImpl struct {
	host string
	auth string
}

// NewDigiqareREST constructor
func NewDigiqareREST(host string, auth string) Digiqare {
	return &digiqareRESTImpl{
		auth: auth,
		host: host,
	}
}

func (d *digiqareRESTImpl) GetMemberById(ctx context.Context, param *MemberParam) (result MemberResponse, code int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "DigiQare:GetMemberById")
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

	if param == nil || strings.TrimSpace(param.ID) == "" {
		return result, http.StatusBadRequest, errors.New("parameter id is required")
	}

	tags := span.Tags()
	tags["param"] = param

	resp.Data = new(MemberResponse)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: d.auth,
	}

	uri = fmt.Sprintf("%s/members/by-user-service-user-id/%s", d.host, param.ID)
	b, statusCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return result, statusCode, err
	}

	json.Unmarshal(b, &resp)
	if member, ok := resp.Data.(*MemberResponse); ok {
		return *member, statusCode, nil
	}

	return result, http.StatusInternalServerError, errors.New("invalid member response type")
}

func (d *digiqareRESTImpl) GetCompanyByCode(ctx context.Context, param *CompanyParam) (result CompanyResponse, code int, err error) {
	span, ctx := tracer.StartTraceWithContext(ctx, "DigiQare:GetCompanyByCode")
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

	if param == nil || strings.TrimSpace(param.Code) == "" {
		return result, http.StatusBadRequest, errors.New("parameter code is required")
	}

	tags := span.Tags()
	tags["param"] = param

	resp.Data = new(CompanyResponse)

	headers := map[string]string{
		echo.HeaderContentType:   echo.MIMEApplicationJSON,
		echo.HeaderAuthorization: d.auth,
	}

	uri = fmt.Sprintf("%s/companies/by-code/%s", d.host, param.Code)
	b, statusCode, err := candiutils.NewHTTPRequest().Do(ctx, http.MethodGet, uri, nil, headers)
	if err != nil {
		return result, statusCode, err
	}

	json.Unmarshal(b, &resp)
	if company, ok := resp.Data.(*CompanyResponse); ok {
		return *company, statusCode, nil
	}

	return result, http.StatusInternalServerError, errors.New("invalid company response type")
}
