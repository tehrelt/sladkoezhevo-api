package models

// Factory model info
// @Description Factory information
type Factory struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	PhoneNumber  string `json:"phoneNumber"`
	Year         int    `json:"year"`
	PropertyType string `json:"propertyType"`
}
