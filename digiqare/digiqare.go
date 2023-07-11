package digiqare

import "context"

// Digiqare abstraction
type Digiqare interface {
	GetMemberById(ctx context.Context, param *MemberParam) (response MemberResponse, code int, err error)
	GetCompanyByCode(ctx context.Context, param *CompanyParam) (response CompanyResponse, code int, err error)
}
