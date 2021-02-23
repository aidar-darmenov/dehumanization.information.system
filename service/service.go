package service

import (
	"github.com/aidar-darmenov/dehumanization.information.system/db"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
)

type Service struct {
	Configuration *model.Configuration
	Manager       *db.Manager
	ChannelString chan model.StringElement
	ChannelErrors chan error
	ChannelFiller chan model.HashedStringElement
}

func NewCoreManager(
	//dbm *db.Manager,
	config *model.Configuration,
) *Service {
	coreManager := Service{
		Configuration: config,
		//Manager:       dbm,
		ChannelString: make(chan model.StringElement),
		ChannelErrors: make(chan error),
		ChannelFiller: make(chan model.HashedStringElement),
	}
	return &coreManager
}
