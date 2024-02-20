package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	 GetAllCities
//		@Summary		Get all cities
//		@Description Get all cities
//		@Produce		json
//		@Success		200	{array}	models.City
//		@Failure		400
//		@Failure		500
//		@Router			/api/v1/cities [get]
func (s *Router) HandlerGetCities(c *fiber.Ctx) error {

	cities, err := s.services.Cities.Get()
	if err != nil {
		return err
	}

	return s.respond(c, cities)
}
