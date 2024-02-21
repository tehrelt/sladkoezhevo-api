package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//	 GetAllPackagingTypes
//		@Summary		Get all packaging methods
//		@Description 	Get all packaging methods
//		@Tags 			Packaging
//		@Produce		json
//		@Success		200	{array}	models.Packaging
//		@Failure		400
//		@Failure		500
//		@Router			/api/v1/packaging [get]
func (s *Router) HandlerPackagingTypes(c *fiber.Ctx) error {

	return s.notimplemented("packaging-handler.HandlerPackagingTypes")

	cities, err := s.services.Cities.Get()
	if err != nil {
		return err
	}

	return s.respond(c, cities)
}

// GetPackagingTypeById
//
//	@Summary		Get packaging by ID
//	@Description 	Get packaging by ID
//	@Tags 			Packaging
//	@Produce		json
//	@Success		200	{array}	models.Packaging
//	@Failure		400
//	@Failure		500
//	@Param 			id path int true "Packaging ID"
//	@Router			/api/v1/packaging/{id} [get]
func (s *Router) HandlerPackagingType(c *fiber.Ctx) error {

	return s.notimplemented("packaging-handler.HandlerPackagingType")

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
