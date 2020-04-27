package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`

}

func NewBadRequestError(message string)*RestErr{
	return &RestErr{
		Message: "Invalid json body",
		Status: http.StatusBadRequest,
		Error: "Bad_request ",


	}
}
