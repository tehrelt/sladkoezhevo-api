package handlers

import (
	"fmt"
	"log/slog"
	_ "sladkoezhevo-api/docs"
	"sladkoezhevo-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type HandlerFunc func(c *fiber.Ctx) error

type Router struct {
	services *services.Services
	logger   *slog.Logger
	App      *fiber.App
}

func New(services *services.Services, logger *slog.Logger) *Router {

	r := &Router{
		services: services,
		logger:   logger,
		App:      fiber.New(),
	}

	r.Configure()

	return r
}

func (s *Router) Configure() {

	s.App.Use(logger.New())

	s.App.Get("/docs/*", swagger.HandlerDefault)

	route := s.App.Group("api/v1")

	route.Get("/ping", s.PingHandler())
	route.Get("/cities", s.HandlerGetCities)
	route.Get("/cities/:id", s.HandlerGetCity)

	s.logger.Info("Routes configured")
}

func (s *Router) Start(port string) error {
	return s.App.Listen(fmt.Sprintf(":%s", port))
}

func (s *Router) PingHandler() HandlerFunc {
	return func(c *fiber.Ctx) error {
		return s.respond(c, "pong")
	}
}

func (s *Router) respond(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(&fiber.Map{
		"data": data,
	})
}

func (s *Router) bad(message string) error {
	return fiber.NewError(500, message)
}

func (s *Router) notfound(message string) error {
	return fiber.NewError(404, message)
}

func (s *Router) error(code int, message string) error {
	return fiber.NewError(code, message)
}
