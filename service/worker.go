package service

import (
	"crypto/sha256"
	"encoding/json"
)

func (s *Service) Worker(channelString chan string) {
	for {
		select {
		case stringToEncrypt := <-s.ChannelString:
			byteArray, err := json.Marshal(stringArray.StringArray[i])
			if err != nil {
				channelErrors <- err
			}
			sha256.Sum256(byteArray)

		}
	}
}
