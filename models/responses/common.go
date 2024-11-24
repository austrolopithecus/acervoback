package responses

type CommonResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
