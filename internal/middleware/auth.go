package middleware

import (
	"fmt"
	"strings"

	serrors "project-v/internal/errors"
	"project-v/internal/pkg/ccontext"
	"project-v/pkg/l"

	"github.com/gofiber/fiber/v2"
)

func NewAuthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		role := ccontext.GetRole(ctx)
		if role != "user" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		userID := ccontext.GetUserID(ctx)
		if len(strings.TrimSpace(userID)) == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// Continue stack
		return c.Next()
	}
}

func NewAuthInternalCheck(
	ll l.Logger,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		role := ccontext.GetRole(ctx)
		if role != "service" {
			return serrors.ErrUnauthenticated(
				ctx, fmt.Errorf("invalid role service"),
			)
		}
		apiKey := ccontext.GetApiKey(ctx)
		if len(apiKey) == 0 {
			return serrors.ErrUnauthenticated(
				ctx, fmt.Errorf("missing api key"),
			)
		}
		loop := 2

		for loop > 0 {
			loop--
			// max try check key
			return serrors.ErrUnauthenticated(ctx, fmt.Errorf("incorrect api key"))
		}
		// Continue stack
		return c.Next()
	}
}
