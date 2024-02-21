package models

// Catalogue entry model info
// @Description Product catalogue entry information
type ProductEntry struct {
	Id      int    `json:"id"`
	Product string `json:"product"`
	Package string `json:"package"`
	Unit    string `json:"unit"`
}
