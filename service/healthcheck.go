package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) HealthCheck(c *gin.Context) {
	t := time.Now().Local()
	c.JSON(200, fmt.Sprintf("ok. Current time:%v", t))
}
