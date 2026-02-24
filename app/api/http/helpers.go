package http

import (
	"context"
	"errors"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/infra/logger"
)

func HandleError(ctx context.Context, errorToHandle error, log logger.Logger) *HttpResponse {
	originalError := application.GetOriginalError(errorToHandle)

	log.ErrorF("Error handled: %s", errorToHandle.Error())

	if errors.As(originalError, &application.NotFoundApplicationError{}) {
		return NotFound(originalError.Error())
	}

	if errors.As(originalError, &application.TimeoutExceededApplicationError{}) {
		return TimeoutExceeded(originalError.Error())
	}

	if errors.As(originalError, &application.IntegrationApplicationError{}) {
		return BadGateway(originalError.Error())
	}

	if errors.As(originalError, &application.BadRequestApplicationError{}) {
		return BadRequest(originalError.Error())
	}

	if errors.As(originalError, &application.ValidationApplicationError{}) {
		return UnprocessableEntity(originalError.Error())
	}

	if errors.As(originalError, &application.ForbiddenApplicationError{}) {
		return Forbidden(originalError.Error())
	}

	return InternalServerError(originalError.Error())
}
