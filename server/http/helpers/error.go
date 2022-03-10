package helpers

import (
	"github.com/diez37/go-packages/log"
	"net/http"
)

type Error struct {
	logger log.Logger
}

func NewError(logger log.Logger) *Error {
	return &Error{logger: logger}
}

func (helper *Error) Error(statusCode int, err error, responseWriter http.ResponseWriter) {
	helper.logger.Error(err)
	responseWriter.WriteHeader(statusCode)
	_, err = responseWriter.Write([]byte(err.Error()))
	if err != nil {
		helper.logger.Error(err)
	}
}
