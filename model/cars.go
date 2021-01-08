package model

import "github.com/go-openapi/strfmt"

type CarsInOut struct {
	LicenseID       string  `json:"license_id"`
	Road            string  `json:"road"`
	Hours           int     `json:"hours"`
	PaymentTypeName string  `json:"payment_type_name"`
	IsFined         bool    `json:"is_fined"`
	FineAmount      float64 `json:"fine_amount"`
}

type Cars struct {
	LicenseID     string       `json:"license_id";gorm:"license_id"`
	Road          string       `json:"road"; gorm:"road"`
	Hours         int          `json:"hours"; gorm:"hours"`
	PaymentTypeID strfmt.UUID4 `json:"payment_type_name"; gorm:"payment_type_id"`
	IsFined       bool         `json:"is_fined"; gorm:"is_fined"`
	FineAmount    float64      `json:"fine_amount"; gorm:"fine_amount"`
}

type PaymentTypes struct {
	ID   strfmt.UUID4
	Name string
}
