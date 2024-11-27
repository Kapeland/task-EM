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
)

type MusicModelManager interface {
	GetLibInfo(ctx context.Context) ([]structs.FullMusicEntry, error)
	GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DeleteSong(ctx context.Context, group string, name string) error
	ChangeSongText(ctx context.Context, group string, newGroup string, name string, newName string) error
	AddSong(ctx context.Context, id int) (structs.TestFull, error)
}

type musicServer struct {
	m MusicModelManager
}

func (s *musicServer) GetLibInfo(c *gin.Context) {
	data, status := s.getLibInfo(c.Request.Context())

	c.JSON(status, data)
}

func (s *musicServer) getLibInfo(ctx context.Context) (GetLibraryContentResp, int) {
	songs, err := s.m.GetLibInfo(ctx)
	if err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return GetLibraryContentResp{}, http.StatusInternalServerError
	}
	libraryContent := make([]FullSongContent, len(songs))

	for i, song := range songs {
		libraryContent[i] = FullSongContent{
			Group:   song.Group,
			Name:    song.Name,
			Text:    song.Text,
			Release: song.Release.Format("2006-01-02"),
			Link:    song.Link,
		}
	}

	return GetLibraryContentResp{
		Library: libraryContent,
	}, http.StatusOK
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
	var sr DeleteSongReq
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status := s.removeSong(c.Request.Context(), sr)

	c.JSON(status, []byte(""))
}
func (s *musicServer) removeSong(ctx context.Context, sr DeleteSongReq) int {
	err := s.m.DeleteSong(ctx, sr.Group, sr.Name)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return http.StatusNotFound
		}
		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (s *musicServer) ChangeSong(c *gin.Context) {
	var sr ChangeSongReq
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status := s.changeSong(c.Request.Context(), sr)

	c.JSON(status, []byte(""))
}
func (s *musicServer) changeSong(ctx context.Context, sr ChangeSongReq) int {
	err := s.m.ChangeSongText(ctx, sr.Group, sr.NewGroup, sr.Name, sr.NewName)
	if err != nil {
		if errors.Is(err, models.ErrConflict) {
			return http.StatusBadRequest
		}
		if errors.Is(err, models.ErrNotFound) {
			return http.StatusNotFound
		}
		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (s *musicServer) AddSong(c *gin.Context) {
	//TODO: Этот метод должен обращаться к стороннему АПИ при добавлении песни, чтобы получить текст и ссылку
	return
}
