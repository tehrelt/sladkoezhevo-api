package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log/slog"
	_ "sladkoezhevo-api/docs"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type HandlerFunc func(c *fiber.Ctx) error

type Router struct {
	services *services.Services
	logger   *slog.Logger
	App      *fiber.App
}

func New(services *services.Services, logger *slog.Logger) *Router {

	app := fiber.New(fiber.Config{
		AppName:       "sladkoezhevo-api",
		CaseSensitive: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).JSON(models.ErrorResponse{
				Message: e.Message,
			})

			return nil
		},
	})

	r := &Router{
		services: services,
		logger:   logger,
		App:      app,
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
	route.Post("/cities", s.HandlerCreateCity)
	route.Put("/cities", s.HandlerUpdateCity)

	route.Get("/packaging/", s.HandlerPackagingTypes)
	route.Get("/packaging/:id", s.HandlerPackagingType)

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
