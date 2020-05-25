package exception

import (
	"fmt"
	"net/http"
)

type Exception struct {
	HttpStatusCode int
	Code           int
	Message        string
}

var (
	MissongRequiredParameterError = Exception{http.StatusBadRequest, 10000, "Missong required parameter: %v"}
)

func FormatException(err Exception, parameters ...interface{}) Exception {
	tmpErr := err
	tmpErr.Message = fmt.Sprintf(tmpErr.Message, parameters...)
	return tmpErr
}
