package customerror

import "fmt"

type DataAlreadyExistsError struct {
	dataType string
}

func (e *DataAlreadyExistsError) Error() string {
	errorMessage := fmt.Sprintf("%s already exists", e.dataType)
	return errorMessage
}

func NewDataAlreadyExistsError(dataType string) *DataAlreadyExistsError {
	return &DataAlreadyExistsError{
		dataType: dataType,
	}
}
