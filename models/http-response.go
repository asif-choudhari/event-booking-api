package models

type Success struct {
	StatusCode int `json:"statusCode"`
	Data       any `json:"data"`
}

type Fail struct {
	StatusCode int `json:"statusCode"`
	// Message    string `json:"message"`
	Message any `json:"errorMessage"`
}
