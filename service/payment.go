package service

import (
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Service) Payment(c *gin.Context) {

	var carInOut model.CarsInOut
	var err error

	err = c.BindJSON(&carInOut)
	if err != nil {
		c.JSON(400, err)
		return
	}

	var paymentType model.PaymentType

	err = s.Manager.DB.First(&paymentType).Where("name = ?", carInOut.PaymentTypeName).Error
	if err != nil {
		c.JSON(400, err)
		return
	}

	var car = model.Cars{
		LicenseID:     carInOut.LicenseID,
		Road:          carInOut.Road,
		Hours:         carInOut.Hours,
		PaymentTypeID: paymentType.ID,
		IsFined:       carInOut.IsFined,
		FineAmount:    carInOut.FineAmount,
	}

	err = s.Manager.DB.Create(&car).Error
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "ok")
}
func (s *Service) Fine(c *gin.Context) {
	var err error
	var car model.Cars
	c.Bind(&car)

	var found bool
	err = s.Manager.DB.First(&car).Where("license_id = ?", car.LicenseID).Error
	if err != gorm.ErrRecordNotFound {
		if err != nil {
			c.JSON(400, err)
			return
		}
	} else {
		found = true
		err = nil
	}

	if found {
		err = s.Manager.DB.Model(&car).Updates(model.Cars{IsFined: car.IsFined, FineAmount: car.FineAmount}).Where("license_id = ?", car.LicenseID).Error
		if err != nil {
			c.JSON(400, err)
			return
		}
	} else {
		err = s.Manager.DB.Create(&car).Error
		if err != nil {
			c.JSON(400, err)
			return
		}
	}

	c.JSON(200, "ok")
}
func (s *Service) Check(c *gin.Context) {
	var car model.Cars
	var err error

	err = c.BindJSON(&car)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = s.Manager.DB.First(&car).Where("license_id = ?", car.LicenseID).Error
	if err != gorm.ErrRecordNotFound {
		if err != nil {
			c.JSON(400, err)
			return
		}
	} else {
		err = nil
		car.IsFined = true
	}

	c.JSON(200, car.IsFined)
}

func (s *Service) Remove(c *gin.Context) {
	licenseID := c.Param("licenseid")
	var car model.Cars

	err := s.Manager.DB.Where("license_id = ?", licenseID).Delete(&car).Error
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, "ok")
}
