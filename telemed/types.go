package telemed

// RegisterUserSendBird data structure
type RequestRegisterUserSendBird struct {
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photoUrl"`
}

type ResponseRegisterUserSendBird struct {
	SbID          string `json:"sbId"`
	SbName        string `json:"sbName"`
	SbPhotoURL    string `json:"sbPhotoUrl"`
	SbAccessToken string `json:"sbAccessToken"`
}

type RestAPIResult struct {
	Errors []ErrRes    `json:"errors"`
	Meta   interface{} `json:"meta"`
	Data   interface{} `json:"data"`
}

type ErrRes struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source string `json:"source"`
}
