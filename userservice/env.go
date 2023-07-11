package userservice

const (

	// Development URL environment
	Development      = "https://api-dev.sehatq.com/v1/user-service"
	AuthorizationDev = "Basic dXNlcnNlcnZpY2U6dXNlcnNlcnZpY2U="

	// Sanity URL environment
	Sanity              = "https://api-sanity.sehatq.com/v1/user-service"
	AuthorizationSanity = "Basic dXNlcnNlcnZpY2VzYW5pdHk6dXNlcnNlcnZpY2VzYW5pdHk="

	// Production URL environment
	Production        = "https://api.sehatq.com/v1/user-service"
	AuthorizationProd = "Basic dXNlcnNlcnZpY2Vwcm9kdWN0aW9uOnVzZXJzZXJ2aWNlcHJvZHVjdGlvbg=="

	// Local URL environment
	Local              = "http://localhost:8000"
	AuthorizationLocal = "Basic dXNlcnNlcnZpY2U6dXNlcnNlcnZpY2U="
)
