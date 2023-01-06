package entities

type Product struct {
	ProductID     int64
	ProductName   string `validate:"required"  label:"ProductName"`
	ProductNumber string `validate:"required"  label:"ProductNumber"`
	ProductPrice  string `validate:"required" label:"ProductPrice"`
}
