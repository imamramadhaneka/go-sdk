package netcore

type AddActivityRequest struct {
	AssetID        string      `json:"asset_id"`
	ActivityName   string      `json:"activity_name"`
	Timestamp      string      `json:"timestamp"`
	Identity       string      `json:"identity"`
	ActivitySource string      `json:"activity_source"`
	ActivityParams interface{} `json:"activity_params"`
}

type AddActivityResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	Error        string `json:"error"`
	SuccessCount string `json:"success_count"`
	FailedCount  string `json:"failed_count"`
}
