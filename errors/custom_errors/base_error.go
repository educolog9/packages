package customerrors

import "fmt"

type BaseErrorInterface interface {
	Error() string
	GetCode() int
	GetMessage() string
	GetInner() error
}

type BaseError struct {
	Msg   string
	Code  int
	Inner error
}

func NewBaseError(msg string, code int, inner error) *BaseError {
	return &BaseError{
		Msg:   msg,
		Code:  code,
		Inner: inner,
	}
}

func (e *BaseError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("Code: %d, Msg: %s, Inner: %v", e.Code, e.Msg, e.Inner)
	}
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func (e *BaseError) GetCode() int {
	return e.Code
}

func (e *BaseError) GetMessage() string {
	return e.Msg
}

func (e *BaseError) GetInner() error {
	return e.Inner
}
