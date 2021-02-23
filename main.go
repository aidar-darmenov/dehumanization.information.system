package main

import (
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/aidar-darmenov/dehumanization.information.system/service"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

func main() {
	_cfg := model.NewConfiguration("config.json")
	_server := server(_cfg)
	_server.Run(":" + strconv.Itoa(_cfg.HTTPPort))
}

func server(_cfg *model.Configuration) *gin.Engine {
	//_db := db.NewDB(_cfg)
	//_dbm := db.NewDBManager(_db)
	//_dbm.AutoMigrate()

	_service := service.NewCoreManager( /*_db,*/ _cfg)

	_service.StartWorkers(_cfg)

	g := gin.Default()
	g.POST("/healthcheck", _service.HealthCheck)
	//g.POST("/payment", _service.Payment)
	//g.PUT("/fine", _service.Fine)
	//g.GET("/check", _service.Check)
	//g.DELETE("/remove/:licenseid", _service.Remove)
	//g.POST("/generate-data", _service.GenerateData)

	g.POST("/encrypt-string-list", _service.EncryptStringList)

	return g
}
