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
	ParseJsonBodyError            = Exception{http.StatusBadRequest, 10001, "Parse Json body failed."}
	QueryDatabaseError            = Exception{http.StatusBadRequest, 10002, "Query Database Failed"}
)

func FormatException(err Exception, parameters ...interface{}) Exception {
	tmpErr := err
	tmpErr.Message = fmt.Sprintf(tmpErr.Message, parameters...)
	return tmpErr
}
