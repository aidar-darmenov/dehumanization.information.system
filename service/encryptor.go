package service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) EncryptStringList(c *gin.Context) {
	var stringArray model.EncryptorStringList

	err := c.Bind(&stringArray)
	if err != nil {
		c.JSON(400, err)
	}

	channelErrors := make(chan error)

	//Starting listener for errors
	for {
		select {
		case errFromChannel := <-channelErrors:
			fmt.Println(errFromChannel) //Here has to be logger used by team
		}
	}

	for i := range stringArray.StringArray {
		go func() {
			byteArray, err := json.Marshal(stringArray.StringArray[i])
			if err != nil {
				channelErrors <- err
			}
			sha256.Sum256(byteArray)
		}()
	}
}
