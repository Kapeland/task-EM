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

	router.GET("/song", implMusic.GetSongText)
	router.GET("/library", implMusic.GetLibInfo)
	router.DELETE("/rmsong", implMusic.RemoveSong)
	router.PUT("/modsong", implMusic.ChangeSong)
	router.POST("/addsong", implMusic.AddSong)

	if err = router.Run(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Service: Launch: router.Run error: %s", err.Error()))
	}
}

//TODO: Проверить, что апи в точности соответствует тому, что происходит
