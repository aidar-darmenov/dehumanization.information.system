package model

import "github.com/go-openapi/strfmt"

type PaymentType struct {
	ID   strfmt.UUID4
	Name string
}

func (PaymentType) TableName() string {
	return "payment_types"
}
