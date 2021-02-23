package service

import (
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

	//Starting listener for errors
	go func() {
		for {
			select {
			case errFromChannel := <-s.ChannelErrors:
				fmt.Println(errFromChannel) //Here has to be the logger used by team
			}
		}
	}()

	//Declaring and making newArray where we will assign new hashed values
	var newArray = make([][32]byte, len(stringArray.StringArray))
	var counter int = 0

	//Starting listener for filling newArray
	go func() {
		var stopFiller bool = false
		for {
			if stopFiller {
				break
			}
			select {
			case newStringElem := <-s.ChannelFiller:
				newArray[newStringElem.Index] = newStringElem.Value
				counter++
				if counter == len(stringArray.StringArray) {
					close(s.ChannelString)
					close(s.ChannelErrors)
					close(s.ChannelFiller)
					c.JSON(200, newArray)
					stopFiller = true
				}
			}
		}
	}()

	//Looping stringArray and sending elements to channel
	for i := range stringArray.StringArray {
		s.ChannelString <- model.StringElement{
			Value: stringArray.StringArray[i],
			Index: i,
		}
	}
}
