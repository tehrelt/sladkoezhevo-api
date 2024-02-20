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

	s.logger.Info("Routes configured")
}

func (s *Server) Start(port string) error {

	s.logger.Info("Starting server", slog.String("port", port))

	if err := s.app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}

func (s *Server) PingHandler() HandlerFunc {
	return func(c *fiber.Ctx) error {
		return c.SendString("pong")
	}
}
