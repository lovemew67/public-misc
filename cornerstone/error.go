package cornerstone

import (
	"fmt"
)

type CodeError struct {
	errorCode    int
	subErrorCode int
	errorMessage string
}

func NewCodeError(ec, sec int, em string) (result CodeError) {
	result = CodeError{
		errorCode:    ec,
		subErrorCode: sec,
		errorMessage: em,
	}
	return
}

func FromNativeError(err error) (result CodeError) {
	result = CodeError{
		errorMessage: err.Error(),
	}
	return
}

func (err *CodeError) Error() (msg string) {
	msg = fmt.Sprintf("%d | %d | %s", err.errorCode, err.subErrorCode, err.errorMessage)
	return
}

func (err *CodeError) ErrorCode() (ec int) {
	ec = err.errorCode
	return
}

func (err *CodeError) SubErrorCode() (sec int) {
	sec = err.subErrorCode
	return
}

func (err *CodeError) ErrorMessage() (em string) {
	em = err.errorMessage
	return
}
