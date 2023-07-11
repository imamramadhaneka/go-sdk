package paymentservice

type RequestRefund struct {
	CoNumber          string  `json:"coNumber"`
	PaymentMethodSlug string  `json:"paymentMethodSlug"`
	MerchantId        int     `json:"merchantId"`
	Amount            float64 `json:"amount"`
	Reason            string  `json:"reason"`
}

type ResponseRefund struct {
	CoNumber          string  `json:"coNumber"`
	PaymentMethodSlug string  `json:"paymentMethodSlug"`
	Amount            float64 `json:"amount"`
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
