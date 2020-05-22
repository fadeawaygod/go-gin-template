package error

import (
	"fmt"
	"net/http"
)

type Error struct {
	HttpStatusCode int
	Code           int
	Message        string
}

var (
	MissongRequiredParameterError = Error{http.StatusBadRequest, 10000, "Missong required parameter: %s"}
)

func FormatError(err *Error, parameters ...interface{}) Error {
	tmpErr := err
	tmpErr.Message = fmt.Sprintf(err.Message, parameters...)
	return *tmpErr
}
