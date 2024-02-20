package handlers

import "github.com/gofiber/fiber/v2"

func (s *Server) CityGet() HandlerFunc {
	return func(c *fiber.Ctx) error {
		cities, err := s.services.Cities.Get()
		if err != nil {
			return err
		}

		return s.respond(c, cities)
	}
}
