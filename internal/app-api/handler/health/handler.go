package health

import (
	"github.com/gofiber/fiber/v2"

	"project-v/pkg/errors"
)

// Controller ...
type Controller struct {
	isReady bool
}

// New ...
func New() Controller {
	return Controller{}
}

// SetReady ...
func (s *Controller) SetReady(b bool) {
	s.isReady = b
}

// Health godoc
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/health [get]
func (s Controller) Health(c *fiber.Ctx) error {
	if s.isReady {
		return c.JSON(
			map[string]string{
				"message": "ok",
			},
		)
	}
	return errors.Error(errors.Unavailable, "Server is shutting down")
}

// Liveness godoc
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/liveness [get]
func (s Controller) Liveness(c *fiber.Ctx) error {
	if s.isReady {
		// todo: check connect db
		return c.JSON(
			map[string]string{
				"message": "ok",
			},
		)
	}
	return errors.Error(errors.Unavailable, "Server is shutting down")
}
