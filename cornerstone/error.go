package cornerstone

import (
	"fmt"
)

type CodeError interface {
	ErrorCode() int
	SubErrorCode() int
	ErrorMessage() string
	Error() string
	FullError() string
}

type CarrierCodeError struct {
	errorCode    int
	subErrorCode int
	errorMessage string
}

func NewCarrierCodeError(ec, sec int, em string) (result CarrierCodeError) {
	result = CarrierCodeError{
		errorCode:    ec,
		subErrorCode: sec,
		errorMessage: em,
	}
	return
}

func FromNativeError(err error) (result CarrierCodeError) {
	result = CarrierCodeError{
		errorMessage: err.Error(),
	}
	return
}

func (err CarrierCodeError) ErrorCode() (ec int) {
	ec = err.errorCode
	return
}

func (err CarrierCodeError) SubErrorCode() (sec int) {
	sec = err.subErrorCode
	return
}

func (err CarrierCodeError) ErrorMessage() (em string) {
	em = err.errorMessage
	return
}

func (err CarrierCodeError) Error() (msg string) {
	msg = fmt.Sprintf("%d | %d | %s", err.errorCode, err.subErrorCode, err.errorMessage)
	return
}

func (err CarrierCodeError) FullError() (msg string) {
	msg = err.Error()
	return
}
