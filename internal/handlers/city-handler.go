package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//	 GetAllCities
//		@Summary		Get all cities
//		@Description 	Get all cities
//		@Tags 			Cities
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

// GetCityById
//
//	@Summary		Get city by ID
//	@Description 	Get city by ID
//	@Tags 			Cities
//	@Produce		json
//	@Success		200	{array}	models.City
//	@Failure		400
//	@Failure		500
//	@Param 			id path int true "City ID"
//	@Router			/api/v1/cities/{id} [get]
func (s *Router) HandlerGetCity(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return s.bad(err.Error())
	}

	city, err := s.services.Cities.GetOne(id)
	if err != nil {
		return s.notfound(fmt.Sprintf("city with id=%d not found", id))
	}

	return s.respond(c, city)
}
