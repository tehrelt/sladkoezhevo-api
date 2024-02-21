package models

import "time"

// Shipment model info
// @Description Shipment information
type Shipment struct {
	Id             int       `json:"id"`
	CatalogueEntry string    `json:"catalogueEntry"`
	Shop           string    `json:"shop"`
	ShipDate       time.Time `json:"shipDate"`
	Quantity       int       `json:"quantity"`
	Cost           float64   `json:"cost"`
}
