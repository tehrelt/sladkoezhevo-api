package models

// Property type model info
// @Description Property type information
// @Description with property type id and name/acronym
type PropertyType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
