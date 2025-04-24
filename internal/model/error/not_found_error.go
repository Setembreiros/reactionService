package customerror

import "fmt"

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	errorMessage := fmt.Sprintf("Data not found")
	return errorMessage
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}
