package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/utils/configer"
	"github.com/Kapeland/task-EM/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

type MusicModelManager interface {
	GetLibInfo(ctx context.Context, group string) ([]structs.FullMusicEntry, error)
	GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DeleteSong(ctx context.Context, group string, name string) error
	ChangeSongText(ctx context.Context, group string, newGroup string, name string, newName string) error
	AddSong(ctx context.Context, fsc structs.FullMusicEntry) error
}

type musicServer struct {
	m MusicModelManager
}

func (s *musicServer) GetLibInfo(c *gin.Context) {
	data, status := s.getLibInfo(c.Request.Context(), c.Query("group"))

	c.JSON(status, data)
}

func (s *musicServer) getLibInfo(ctx context.Context, group string) (GetLibraryContentResp, int) {
	songs, err := s.m.GetLibInfo(ctx, group)
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
	var sr AddSongReq
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	data, status := s.getSongFromRemote(sr)
	if status != http.StatusOK {
		c.JSON(status, []byte(""))
		return
	}

	tmpTime, err := time.Parse("02.01.2006", data.ReleaseDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, []byte(""))
		return
	}

	status = s.addSong(c.Request.Context(), structs.FullMusicEntry{
		Group:   sr.Group,
		Name:    sr.Name,
		Text:    data.Text,
		Release: tmpTime,
		Link:    data.Link,
	})

	c.JSON(status, []byte(""))
}

func (s *musicServer) getSongFromRemote(sr AddSongReq) (GetSongInfoResp, int) {
	cfg, err := configer.GetConfig()
	if err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Service: Launch: configer.GetConfig error: %s", err.Error()))
		return GetSongInfoResp{}, http.StatusInternalServerError
	}

	params := url.Values{}
	params.Add("group", sr.Group)
	params.Add("song", sr.Name)

	remoteURL := fmt.Sprintf("%s://%s:%d/info?%s", cfg.RmServer.Protocol, cfg.RmServer.Host, cfg.RmServer.Port, params.Encode())
	resp, err := http.Get(remoteURL)
	if err != nil {
		logger.Log(logger.InfoPrefix, fmt.Sprintf("Music_server: getSongFromRemote: Get url: %s", remoteURL))
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Music_server: getSongFromRemote: Get error: %s", err.Error()))
		return GetSongInfoResp{}, http.StatusInternalServerError
	}
	if resp.StatusCode != http.StatusOK {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Music_server: getSongFromRemote: StatusCode from remote: %d", resp.StatusCode))
		return GetSongInfoResp{}, http.StatusInternalServerError
	}

	defer resp.Body.Close()

	var remoteSongResp GetSongInfoResp

	if err := json.NewDecoder(resp.Body).Decode(&remoteSongResp); err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("Music_server: getSongFromRemote: Decode error: %s", err.Error()))
		return GetSongInfoResp{}, http.StatusInternalServerError
	}

	return remoteSongResp, http.StatusOK
}

func (s *musicServer) addSong(ctx context.Context, fsc structs.FullMusicEntry) int {
	err := s.m.AddSong(ctx, fsc)
	if err != nil {
		if errors.Is(err, models.ErrConflict) {
			return http.StatusBadRequest
		}

		logger.Log(logger.ErrPrefix, fmt.Sprintf(err.Error()))
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
