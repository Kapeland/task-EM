package storage

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/storage/repository"

	"github.com/pkg/errors"
)

type MusicRepo interface {
	GetAllMusic(ctx context.Context, group string) ([]structs.FullMusicEntry, error)
	GetMusicText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DelSong(ctx context.Context, group string, name string) error
	PutSong(ctx context.Context, group string, newGroup string, name string, newName string) error
	PostSong(ctx context.Context, fsc structs.FullMusicEntry) error
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

func (m *MusicStorage) GetAllMusic(ctx context.Context, group string) ([]structs.FullMusicEntry, error) {
	songs, err := m.musicRepo.GetAllMusic(ctx, group)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (m *MusicStorage) DeleteSong(ctx context.Context, group string, name string) error {
	err := m.musicRepo.DelSong(ctx, group, name)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return models.ErrNotFound
		}
		return err
	}
	return nil
}

func (m *MusicStorage) AddSong(ctx context.Context, fsc structs.FullMusicEntry) error {
	err := m.musicRepo.PostSong(ctx, fsc)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateKey) {
			return models.ErrConflict
		}

		return err
	}

	return nil
}
func (m *MusicStorage) ChangeSongText(ctx context.Context, group string, newGroup string, name string, newName string) error {
	err := m.musicRepo.PutSong(ctx, group, newGroup, name, newName)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateKey) {
			return models.ErrConflict
		}
		if errors.Is(err, repository.ErrObjectNotFound) {
			return models.ErrNotFound
		}

		return err
	}
	return nil
}
