package services

import (
	"context"
	"fmt"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type MusicModelManager interface {
	GetLibInfo(ctx context.Context, id int) (structs.TestFull, error)
	GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DeleteSong(ctx context.Context, id int) (structs.TestFull, error)
	ChangeSongText(ctx context.Context, id int) (structs.TestFull, error)
	AddSong(ctx context.Context, id int) (structs.TestFull, error)
}

type musicServer struct {
	m MusicModelManager
}

func (s *musicServer) GetLibInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	data, status := s.getLibInfo(c.Request.Context(), id)

	c.JSON(status, data)
}

func (s *musicServer) getLibInfo(ctx context.Context, id int) ([]byte, int) {
	_, err := s.m.GetLibInfo(ctx, id)
	if err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return nil, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *musicServer) GetSongText(c *gin.Context) {
	var sr GetSongTextReq
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	data, status := s.getSongText(c.Request.Context(), sr)

	c.JSON(status, data)
}

func (s *musicServer) getSongText(ctx context.Context, sr GetSongTextReq) (GetSongTextResp, int) {
	song, err := s.m.GetSongText(ctx, sr.Group, sr.Name)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return GetSongTextResp{}, http.StatusNotFound
		}

		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return GetSongTextResp{}, http.StatusInternalServerError
	}

	return GetSongTextResp{
		Name: song.Name,
		Text: song.Text,
	}, http.StatusOK
}

func (s *musicServer) RemoveSong(c *gin.Context) {
	return
}

func (s *musicServer) ChangeSong(c *gin.Context) {
	return
}

func (s *musicServer) AddSong(c *gin.Context) {
	//TODO: Этот метод должен обращаться к стороннему АПИ при добавлении песни, чтобы получить текст и ссылку
	return
}
