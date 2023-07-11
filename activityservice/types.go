package activityservice

import "go.mongodb.org/mongo-driver/bson/primitive"

// RequestInsertLog data structure
type RequestInsertLog struct {
	Service  string      `json:"service"`
	Module   string      `json:"module"`
	Action   string      `json:"action"`
	Target   string      `json:"target"`
	Logs     interface{} `json:"logs"`
	IsDelete bool        `json:"isDelete"`
	User     Users       `json:"user"`
}

type ResponseLog struct {
	ID        primitive.ObjectID `json:"id"`
	Service   string             `json:"service"`
	Module    string             `json:"module"`
	Action    string             `json:"action"`
	Target    string             `json:"target"`
	Logs      interface{}        `json:"logs"`
	IsDelete  bool               `json:"isDelete"`
	User      Users              `json:"user"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
}

type Users struct {
	UUID    string `json:"uuid,omitempty" bson:"uuid"`
	Name    string `json:"name,omitempty" bson:"name"`
	Email   string `json:"email,omitempty" bson:"email"`
	Role    string `json:"role,omitempty" bson:"role"`
	Channel string `json:"channel,omitempty" bson:"channel"`
}

type LogParam struct {
	ID                string `json:"id"`
	Service           string `json:"service"`
	Module            string `json:"module"`
	Action            string `json:"action"`
	Target            string `json:"target"`
	IsDelete          string `json:"isDelete"`
	PaymentMethodSlug string `json:"paymentMethodSlug"`
	Page              int    `json:"page"`
	PerPage           int    `json:"perPage"`
	OrderBy           string `json:"orderBy"`
	Sort              string `json:"sort"`
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
