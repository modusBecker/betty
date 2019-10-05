package structures

// StartContainerRequest represents the request body for the start_container process
type StartContainerRequest struct {
	ID          string   `json:"id"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Network     string   `json:"network"`
	Environment []string `json:"environment"`
}

// StartContainerResponse represents the response body for the start_container process
type StartContainerResponse struct {
	ID       string   `json:"id"`
	Warnings []string `json:"warnings"`
}

// ErrorResponse represents the structure of an error API response
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
