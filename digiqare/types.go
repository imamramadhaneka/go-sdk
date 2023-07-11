package digiqare

// MemberParam data structure
type MemberParam struct {
	ID string `json:"id"`
}

type CompanyParam struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type MemberResponse struct {
	ID int `json:"id"`
}

type CompanyResponse struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RestAPIResult struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}
