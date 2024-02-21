package models

// Product model info
// @Description Product information
type Product struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Factory string `json:"factory"`
}
