package application

import (
	"fmt"
	"runtime"
	"strings"

	_errors "github.com/go-errors/errors"
)

var (
	NotFoundApplicationErrorType        = "NotFoundApplicationError"
	TimeoutExceededApplicationErrorType = "TimeoutExceededApplicationError"
	IntegrationApplicationErrorType     = "IntegrationApplicationError"
	ValidationApplicationErrorType      = "ValidationApplicationError"
	BadRequestApplicationErrorType      = "BadRequestApplicationError"
	ForbiddenApplicationErrorType       = "ForbiddenApplicationError"
)

type NotFoundApplicationError struct {
	error
}
type TimeoutExceededApplicationError struct {
	error
}

type IntegrationApplicationError struct {
	error
}

type ValidationApplicationError struct {
	error
}

type ForbiddenApplicationError struct {
	error
}

type BadRequestApplicationError struct {
	error
}

type errorWrapper interface {
	Error() string
	GetOriginalError() error
}

type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

func (err WrappedError) Error() string {
	if len(err.messages) > 0 {
		retVal := fmt.Sprintf("%s: ", err.path)

		for _, message := range err.messages {
			retVal += message + "; "
		}

		return fmt.Sprintf("%s => %v", retVal, err.originalError)
	}

	return fmt.Sprintf("%s => %v", err.path, err.originalError)
}

func (err WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		originalError, ok := (err.originalError).(errorWrapper)
		if ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}

func Wrap(err error, messages ...string) error {
	return &WrappedError{
		originalError: err,
		path:          caller(),
		messages:      messages,
	}
}

func GetOriginalError(err error) error {

	wrappedErr, ok := err.(errorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}

func caller() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	funcRef := runtime.FuncForPC(pc[0])

	pathArr := strings.Split(funcRef.Name(), "/")

	return pathArr[len(pathArr)-1]
}

func NewIntegrationApplicationError(err error) error {
	return newApplicationError(IntegrationApplicationErrorType, err)
}

func NewTimeoutExceededApplicationError(err error) error {
	return newApplicationError(TimeoutExceededApplicationErrorType, err)
}

func NewNotFoundApplicationError(err error) error {
	return newApplicationError(NotFoundApplicationErrorType, err)
}

func NewValidationApplicationError(err error) error {
	return newApplicationError(ValidationApplicationErrorType, err)
}

func NewBadRequestApplicationError(err error) error {
	return newApplicationError(BadRequestApplicationErrorType, err)
}

func NewForbiddenApplicationError(err error) error {
	return newApplicationError(ForbiddenApplicationErrorType, err)
}

func newApplicationError(errType string, err error) error {
	if err == nil {
		return err
	}
	switch errType {
	case NotFoundApplicationErrorType:
		return _errors.Wrap(NotFoundApplicationError{err}, 1)
	case TimeoutExceededApplicationErrorType:
		return _errors.Wrap(TimeoutExceededApplicationError{err}, 1)
	case IntegrationApplicationErrorType:
		return _errors.Wrap(IntegrationApplicationError{err}, 1)
	case ValidationApplicationErrorType:
		return _errors.Wrap(ValidationApplicationError{err}, 1)
	case BadRequestApplicationErrorType:
		return _errors.Wrap(BadRequestApplicationError{err}, 1)
	case ForbiddenApplicationErrorType:
		return _errors.Wrap(ForbiddenApplicationError{err}, 1)
	default:
		return _errors.Wrap(err, 1)
	}
}
