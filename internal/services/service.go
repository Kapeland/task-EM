package services

import (
	"fmt"
	"github.com/Kapeland/task-EM/internal/utils/configer"
	"github.com/Kapeland/task-EM/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Service struct {
	mm MusicModelManager
}

func NewService(mm MusicModelManager) Service {
	return Service{mm: mm}
}

func (s Service) Launch() {
	implMusic := musicServer{s.mm}
	cfg, err := configer.GetConfig()
	if err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Service: Launch: configer.GetConfig error: %s", err.Error()))
		return
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/info", implMusic.GetLibInfo)

	if err = router.Run(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Service: Launch: router.Run error: %s", err.Error()))
	}
}
