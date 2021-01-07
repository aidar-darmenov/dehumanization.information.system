package model

import "github.com/go-openapi/strfmt"

type Cars struct {
	LicenseID   string
	Road        string
	Hours       int
	PaymentType string
	IsFined     bool
	FineAmount  float64
}

type PaymentTypes struct {
	ID   strfmt.UUID4
	Name string
}
