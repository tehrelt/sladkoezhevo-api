package models

// District model info
// @Description District information
// @Description with district id, name and city
type District struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}
