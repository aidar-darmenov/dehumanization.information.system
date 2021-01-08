package service

import (
	"fmt"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/gin-gonic/gin"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var roads = []string{
	"Ohaio",
	"Almaty",
	"Cambridge",
	"1-B",
	"1-A",
}

func (s *Service) GenerateData(c *gin.Context) {

	for i := 0; i < 200; i++ {
		var car = model.Cars{
			LicenseID:     generateLicense(10),
			Road:          roads[rand.Intn(4-0)],
			Hours:         rand.Intn(24 - 0),
			PaymentTypeID: "c57f95b8-e562-4b7e-9deb-483975d2ed5c",
			IsFined:       false,
			FineAmount:    0,
		}
		err := s.Manager.DB.Create(&car).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	c.JSON(200, "ok")
}

func generateLicense(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
