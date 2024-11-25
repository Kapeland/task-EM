package app

import (
	"context"
	"fmt"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/services"
	"github.com/Kapeland/task-EM/internal/storage"
	"github.com/Kapeland/task-EM/internal/storage/file-storage/file_provider"
	"github.com/Kapeland/task-EM/internal/storage/file-storage/music"
	musicPostgres "github.com/Kapeland/task-EM/internal/storage/repository/postgresql/music"
	"github.com/Kapeland/task-EM/internal/utils/logger"
)

func Start() error {
	ctx := context.Background()
	dbStor, err := storage.NewDbStorage(ctx)
	if err != nil {
		logger.Log(logger.ErrPrefix, fmt.Sprintf("App: Start: NewDbStorage: %s", err.Error()))
		return err
	}
	defer dbStor.Close(ctx)

	musicRepo := musicPostgres.New(dbStor.DB)

	f := file_provider.NewFileProvider()

	mp := music.NewRepository(f)

	musicStorage := storage.NewMusicStorage(musicRepo, mp)

	mmdl := models.NewModelMusic(&musicStorage)

	serv := services.NewService(&mmdl)

	serv.Launch()

	return nil
}
