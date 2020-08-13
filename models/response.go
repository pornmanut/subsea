package models

type TokenResponse struct {
	Token string `json:"token"`
}
type ValidateErrorsResponse struct {
	Errors []ValidateErrorResponse `json:"errors"`
}

type ValidateErrorResponse struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
}

// ErrorResponse reponse message for error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

// SuccessCreated reponse message when success put into database
type SuccessCreated struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
