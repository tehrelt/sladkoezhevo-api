package handlers

import (
	"fmt"
	"log/slog"
	"sladkoezhevo-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type HandlerFunc func(c *fiber.Ctx) error

type Server struct {
	app      *fiber.App
	services *services.Services
	logger   *slog.Logger
}

func New(services *services.Services, logger *slog.Logger) *Server {

	return &Server{
		app:      fiber.New(),
		services: services,
		logger:   logger,
	}
}

func (s *Server) Configure() {
	s.app.Get("/ping", s.PingHandler())

	s.app.Get("/cities", s.CityGet())

	s.logger.Info("Routes configured")
}

func (s *Server) Start(port string) error {

	s.logger.Info("Starting server", slog.String("port", port))

	return s.app.Listen(fmt.Sprintf(":%s", port))
}

func (s *Server) PingHandler() HandlerFunc {
	return func(c *fiber.Ctx) error {
		return s.respond(c, "pong")
	}
}

func (s *Server) respond(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(&fiber.Map{
		"data": data,
	})
}
