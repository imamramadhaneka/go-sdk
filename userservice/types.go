package userservice

import (
	"encoding/json"
)

// VerifyRequest
type VerifyRequest struct {
	Token string `json:"token"`
}

// VerifyResponse
type VerifyResponse struct {
	UUID         string   `json:"uuid,omitempty"`
	Application  []string `json:"application,omitempty"`
	Fullname     string   `json:"fullName"`
	Email        string   `json:"email"`
	PhotoUrl     string   `json:"photoUrl"`
	Token        string   `json:"token"`
	RefreshToken string   `json:"refreshToken"`
}

// FilterUser model
type FilterUser struct {
	UUID        string   `json:"uuid"`
	Fullname    string   `json:"fullname"`
	Email       string   `json:"email"`
	Search      string   `json:"search"`
	Status      string   `json:"status"`
	Application []string `json:"application"`
	Page        int      `json:"page"`
	PerPage     int      `json:"perPage"`
	OrderBy     string   `json:"orderBy"`
	Sort        string   `json:"sort"`
}

type UserResponse struct {
	UUID        string          `json:"uuid"`
	Fullname    string          `json:"fullname"`
	Email       string          `json:"email"`
	Status      string          `json:"status"`
	Application json.RawMessage `json:"application,omitempty"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
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
