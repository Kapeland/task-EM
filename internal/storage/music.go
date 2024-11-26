package storage

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/storage/repository"

	"github.com/pkg/errors"
)

type MusicRepo interface {
	GetAllMusic(ctx context.Context, id int) (structs.TestFull, error)
	GetMusicText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DelSong(ctx context.Context, id int) (structs.TestFull, error)
	PutSong(ctx context.Context, id int) (structs.TestFull, error)
	PostSong(ctx context.Context, id int) (structs.TestFull, error)
}

type MusicStorage struct {
	musicRepo MusicRepo
	mp        MusicProvider
}

type MusicProvider interface {
	GetMusicByte(music string) ([]byte, error)
}

func NewMusicStorage(musicRepo MusicRepo, mp MusicProvider) MusicStorage {
	return MusicStorage{musicRepo: musicRepo, mp: mp}
}

func (m *MusicStorage) GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error) {
	song, err := m.musicRepo.GetMusicText(ctx, group, name)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return structs.MusicEntry{}, models.ErrNotFound
		}

		return structs.MusicEntry{}, err
	}
	return song, nil
}

func (m *MusicStorage) GetAllMusic(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}

func (m *MusicStorage) DeleteSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *MusicStorage) AddSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *MusicStorage) ChangeSongText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
