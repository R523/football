package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Rotate struct {
	Channel chan<- int
}

func (r *Rotate) Get(c *fiber.Ctx) error {
	angle, _ := c.ParamsInt("angle", 0)

	r.Channel <- angle

	// nolint: wrapcheck
	return c.SendStatus(http.StatusNoContent)
}

func (r *Rotate) Register(g fiber.Router) {
	g.Get("/rotate/:angle", r.Get)
}
