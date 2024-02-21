package models

// Shop model info
// @Description Shop information
type Shop struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	District       string `json:"district"`
	PhoneNumber    string `json:"phoneNumber"`
	Year           int    `json:"year"`
	EmployeesCount string `json:"employeesCount"`
}
