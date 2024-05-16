package middleware

import (
	"context"
	"sync"
	"unsafe"

	"project-v/internal/constant"
	"project-v/pkg/l"
	tracer "project-v/pkg/trace"

	"github.com/gofiber/fiber/v2"
)

func getString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// NewLogging creates a new middleware handler
func NewLogging(ll l.Logger) fiber.Handler {
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		once.Do(
			func() {
				errHandler = c.App().Config().ErrorHandler
			},
		)
		ctx := c.UserContext()
		ctx, span := tracer.StartSpan(ctx, "MiddlewareLogging "+c.Path())
		defer span.End()

		for _, item := range []string{
			fiber.HeaderXRequestID,
			constant.XGapoRole,
			constant.XGapoApiKey,
			constant.XGapoUserId,
			constant.XGapoWorkspaceId,
			constant.XGapoLang,
			constant.HeaderTimezoneOffset,
			fiber.HeaderUserAgent,
		} {
			val := getString(c.Context().Request.Header.Peek(item))
			ctx = context.WithValue(ctx, item, val)

			if item == constant.XGapoApiKey {
				apiKeyLen := len(val)
				if apiKeyLen > 3 {
					val = val[0:apiKeyLen-3] + "***"
				}
			}
			tracer.SetAttribute(span, "headers."+item, val)
		}

		c.Locals(constant.SpanKey, ctx)
		c.SetUserContext(ctx)

		defer func() {
			tracer.SetAttribute(span, "code", c.Response().StatusCode())
			tracer.SetAttribute(
				span, "user-agent",
				getString(c.Context().Request.Header.Peek(fiber.HeaderUserAgent)),
			)
		}()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		return nil
	}
}
