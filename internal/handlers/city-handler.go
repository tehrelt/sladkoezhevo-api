package handlers

import (
	"errors"
	"fmt"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateCity
//
//	@Summary		Create city
//	@Description 	Create city
//	@Tags 			Cities
//	@Param			city body requests.CityCreate true "City model"
//	@Produce		json
//	@Success		201 "Successfully created"
//	@Failure		500
//	@Router			/api/v1/cities [post]
func (s *Router) HandlerCreateCity(c *fiber.Ctx) error {

	var city models.City

	if err := c.BodyParser(&city); err != nil {
		return s.internal(err.Error())
	}

	if err := s.services.Cities.Create(&city); err != nil {
		return s.internal(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetAllCities
//
//	@Summary		Get all cities
//	@Description 	Get all cities
//	@Tags 			Cities
//	@Produce		json
//	@Success		200	{array}	models.City
//	@Failure		404 "Cities didn't found"
//	@Failure		500
//	@Router			/api/v1/cities [get]
func (s *Router) HandlerGetCities(c *fiber.Ctx) error {

	cities, err := s.services.Cities.Get()
	if err != nil {
		return s.notfound("no cities")
	}

	return s.respond(c, cities)
}

// GetCityById
//
//	@Summary		Get city by ID
//	@Description 	Get city by ID
//	@Tags 			Cities
//	@Produce		json
//	@Success		200	{object}	models.City
//	@Failure		400
//	@Failure		404 "City didn't find in database"
//	@Failure		500
//	@Param 			id path int true "City ID"
//	@Router			/api/v1/cities/{id} [get]
func (s *Router) HandlerGetCity(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return s.bad(fmt.Sprintf("city id must be positive integer"))
		}
		return s.internal(fmt.Sprintf("%s [@HandlerGetCity]", err.Error()))
	}

	city, err := s.services.Cities.GetOne(id)
	if err != nil {
		if errors.Is(err, storage.ErrRecordNotFound) {
			return s.notfound(fmt.Sprintf("city with id=%d not found", id))
		}
		return s.internal(fmt.Sprintf("%s [@HandlerGetCity]", err.Error()))
	}

	return s.respond(c, city)
}

// UpdateCity
//
//	@Summary		Update city
//	@Description 	Updates city
//	@Tags 			Cities
//	@Param			city body models.City true "City model"
//	@Produce		json
//	@Success		200 "Successfully updated"
//	@Failure		404 "City didn't found in database"
//	@Failure		500
//	@Router			/api/v1/cities [put]
func (s *Router) HandlerUpdateCity(c *fiber.Ctx) error {

	var city models.City

	if err := c.BodyParser(&city); err != nil {
		return err
	}

	if err := s.services.Cities.Update(&city); err != nil {
		if errors.Is(err, storage.ErrRecordNotFound) {
			return s.notfound(fmt.Sprintf("city with id=%d not found", city.Id))
		}
		return s.internal(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
