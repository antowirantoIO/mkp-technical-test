package utils

import (
	"mkp-boarding-test/internal/model"

	"github.com/gofiber/fiber/v2"
)

// SuccessResponse creates a standardized success response
func SuccessResponse[T any](message string, data T) model.WebResponse[T] {
	return model.WebResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// SuccessResponseWithMeta creates a standardized success response with pagination metadata
func SuccessResponseWithMeta[T any](message string, data T, meta *model.PageMetadata) model.WebResponse[T] {
	return model.WebResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

// ErrorResponse creates a standardized error response
func ErrorResponse(message string, errors string) model.WebResponse[interface{}] {
	return model.WebResponse[interface{}]{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}

// CreatePaginationMeta creates pagination metadata with proper calculations
func CreatePaginationMeta(page, size int, total int64) *model.PageMetadata {
	lastPage := (total + int64(size) - 1) / int64(size)
	if lastPage == 0 {
		lastPage = 1
	}

	from := (page-1)*size + 1
	to := page * size
	if int64(to) > total {
		to = int(total)
	}
	if total == 0 {
		from = 0
		to = 0
	}

	return &model.PageMetadata{
		CurrentPage: page,
		PerPage:     size,
		Total:       total,
		LastPage:    lastPage,
		From:        from,
		To:          to,
	}
}

// SendSuccessResponse sends a success response with proper HTTP status
func SendSuccessResponse[T any](ctx *fiber.Ctx, message string, data T) error {
	return ctx.Status(fiber.StatusOK).JSON(SuccessResponse(message, data))
}

// SendSuccessResponseWithMeta sends a success response with pagination metadata
func SendSuccessResponseWithMeta[T any](ctx *fiber.Ctx, message string, data T, meta *model.PageMetadata) error {
	return ctx.Status(fiber.StatusOK).JSON(SuccessResponseWithMeta(message, data, meta))
}

// SendCreatedResponse sends a created response (201)
func SendCreatedResponse[T any](ctx *fiber.Ctx, message string, data T) error {
	return ctx.Status(fiber.StatusCreated).JSON(SuccessResponse(message, data))
}

// SendErrorResponse sends an error response with appropriate HTTP status
func SendErrorResponse(ctx *fiber.Ctx, statusCode int, message string, errors string) error {
	return ctx.Status(statusCode).JSON(ErrorResponse(message, errors))
}

// SendBadRequestResponse sends a 400 Bad Request response
func SendBadRequestResponse(ctx *fiber.Ctx, message string, errors string) error {
	return SendErrorResponse(ctx, fiber.StatusBadRequest, message, errors)
}

// SendUnauthorizedResponse sends a 401 Unauthorized response
func SendUnauthorizedResponse(ctx *fiber.Ctx, message string) error {
	return SendErrorResponse(ctx, fiber.StatusUnauthorized, message, "")
}

// SendForbiddenResponse sends a 403 Forbidden response
func SendForbiddenResponse(ctx *fiber.Ctx, message string) error {
	return SendErrorResponse(ctx, fiber.StatusForbidden, message, "")
}

// SendNotFoundResponse sends a 404 Not Found response
func SendNotFoundResponse(ctx *fiber.Ctx, message string) error {
	return SendErrorResponse(ctx, fiber.StatusNotFound, message, "")
}

// SendInternalServerErrorResponse sends a 500 Internal Server Error response
func SendInternalServerErrorResponse(ctx *fiber.Ctx, message string) error {
	return SendErrorResponse(ctx, fiber.StatusInternalServerError, message, "")
}
