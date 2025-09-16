package response

type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Errors  string `json:"errors"`
}
