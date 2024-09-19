package web

type CommonResponse struct {
	ApiVersion string `json:"api_version,omitempty"`
	Status     string `json:"status,omitempty"`
	IsSuccess  bool   `json:"is_success,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}
