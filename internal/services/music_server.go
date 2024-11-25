package services

import (
	"context"
	"fmt"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MusicModelManager interface {
	GetLibInfo(ctx context.Context, id int) (structs.TestFull, error)
	GetSongText(ctx context.Context, id int) (structs.TestFull, error)
	DeleteSong(ctx context.Context, id int) (structs.TestFull, error)
	ChangeSongText(ctx context.Context, id int) (structs.TestFull, error)
	AddSong(ctx context.Context, id int) (structs.TestFull, error)
}

type musicServer struct {
	m MusicModelManager
}

const categoryParam = "category"
const limitParam = "limit"
const offsetParam = "offset"

//func (s *musicServer) GetAllTests(w http.ResponseWriter, req *http.Request) {
//	category := req.URL.Query().Get(categoryParam)
//	limit := req.URL.Query().Get(limitParam)
//	offset := req.URL.Query().Get(offsetParam)
//	par := getAllTestsParams{}
//	par.fromStrings(category, limit, offset)
//	data, status := s.getAllTests(req.Context(), par)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	_, err := w.Write(data)
//	if err != nil {
//		return
//	}
//}

type getAllTestsParams struct {
	Category string
	Limit    int
	Offset   int
}

func (p *getAllTestsParams) fromStrings(category, limit, offset string) {
	p.Category = category
	p.Limit, _ = strconv.Atoi(limit)
	p.Offset, _ = strconv.Atoi(offset)
}

//func (s *musicServer) getAllTests(ctx context.Context, par getAllTestsParams) ([]byte, int) {
//	t, err := s.m.GetLibInfo(ctx, par.Category, par.Limit, par.Offset)
//	if err != nil {
//		logger.Log(logger.ErrPrefix, fmt.Sprintf("GetMusic: %v", err))
//		return nil, http.StatusInternalServerError
//	}
//
//	tests := make([]GetAllTestsRespTest, 0)
//	for _, t := range t {
//		pic64 := base64.StdEncoding.EncodeToString(t.Picture)
//		tests = append(tests, GetAllTestsRespTest{
//			Id:          t.ID,
//			Name:        t.Name,
//			Description: t.Description,
//			Category:    t.Category,
//			DiffLevel:   t.DiffLevel,
//			Picture:     pic64,
//		})
//	}
//
//	articleJSON, _ := json.Marshal(
//		GetAllTestsResp{
//			Tests: tests,
//		},
//	)
//	return articleJSON, http.StatusOK
//}

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
	return
}

func (s *musicServer) getSongText(ctx context.Context, id int) ([]byte, int) {
	return nil, 0
}

//TODO: доделать необходимые методы, чтобы дёргать их из сервера
