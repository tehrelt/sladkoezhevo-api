package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Router) respond(c *fiber.Ctx, code int, data interface{}) error {
	return c.Status(code).JSON(&fiber.Map{
		"data": data,
	})
}

func (s *Router) bad(message string) error {
	return fiber.NewError(400, message)
}
func (s *Router) notfound(message string) error {
	return fiber.NewError(404, message)
}
func (s *Router) internal(message string) error {
	return fiber.NewError(500, message)
}
func (s *Router) notimplemented(fn string) error {
	return fiber.NewError(500, fmt.Sprintf("%s: not implemented", fn))
}
