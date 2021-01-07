package service

import (
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Service) Payment(c *gin.Context) {
	var car model.Cars
	var err error

	err = c.BindJSON(&car)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = s.Manager.DB.Create(&car).Error
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "ok")
}
func (s *Service) Fine(c *gin.Context) {
	var car model.Cars
	err := c.BindJSON(&car)
	if err != nil {
		c.JSON(400, err)
		return
	}

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
		err = s.Manager.DB.Save(&car).Where("license_id = ?", car.LicenseID).Error
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
