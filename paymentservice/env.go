package paymentservice

const (

	// Development URL environment
	Development      = "https://api-dev.sehatq.com/v1/payment-service"
	AuthorizationDev = "Basic cGF5bWVudHNlcnZpY2U6cGF5bWVudHNlcnZpY2U="

	// Sanity URL environment
	Sanity              = "https://api-sanity.sehatq.com/v1/payment-service"
	AuthorizationSanity = "Basic cGF5bWVudHNlcnZpY2VzYW5pdHk6cGF5bWVudHNlcnZpY2VzYW5pdHk="

	// Production URL environment
	Production        = "https://api.sehatq.com/v1/payment-service"
	AuthorizationProd = "Basic cGF5bWVudHNlcnZpY2Vwcm9kOnBheW1lbnRzZXJ2aWNlcHJvZA=="

	// Local URL environment
	Local              = "http://localhost:8000"
	AuthorizationLocal = "Basic cGF5bWVudHNlcnZpY2U6cGF5bWVudHNlcnZpY2U="
)
