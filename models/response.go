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
	Error string `json:"error"`
}
