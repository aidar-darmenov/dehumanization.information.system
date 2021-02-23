package service

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
)

func (s *Service) StartWorkers(cfg *model.Configuration) {
	for i := 0; i < cfg.WorkersNumber; i++ {
		go s.Worker()
	}
}

func (s *Service) Worker() {
	for {
		select {
		case stringElementToEncrypt := <-s.ChannelString:
			byteArray, err := json.Marshal(stringElementToEncrypt.Value)
			if err != nil {
				//logger used here
			}
			s.ChannelFiller <- model.HashedStringElement{
				Value: sha256.Sum256(byteArray),
				Index: stringElementToEncrypt.Index,
			}
		}
	}
}
