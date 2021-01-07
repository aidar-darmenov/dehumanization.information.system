package service

import (
	"github.com/aidar-darmenov/dehumanization.information.system/db"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
)

type Service struct {
	Configuration *model.Configuration
	Manager       *db.Manager
}

func NewCoreManager(
	dbm *db.Manager,
	config *model.Configuration,
) *Service {
	coreManager := Service{
		Configuration: config,
		Manager:       dbm,
	}
	return &coreManager
}
