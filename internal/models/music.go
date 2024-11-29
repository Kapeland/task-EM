package models

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/pkg/errors"
)

type MusicStorager interface {
	GetAllMusic(ctx context.Context, group string) ([]structs.FullMusicEntry, error)
	GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error)
	DeleteSong(ctx context.Context, group string, name string) error
	AddSong(ctx context.Context, fsc structs.FullMusicEntry) error
	ChangeSongText(ctx context.Context, group string, newGroup string, name string, newName string) error
}

const GroupFilter = ""

func (m *ModelMusic) GetLibInfo(ctx context.Context, group string) ([]structs.FullMusicEntry, error) {

	songs, err := m.ms.GetAllMusic(ctx, group)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
func (m *ModelMusic) GetSongText(ctx context.Context, group string, name string) (structs.MusicEntry, error) {
	song, err := m.ms.GetSongText(ctx, group, name)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return structs.MusicEntry{}, ErrNotFound
		}

		return structs.MusicEntry{}, err
	}

	return song, nil
}
func (m *ModelMusic) DeleteSong(ctx context.Context, group string, name string) error {
	err := m.ms.DeleteSong(ctx, group, name)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
func (m *ModelMusic) ChangeSongText(ctx context.Context, group string, newGroup string, name string, newName string) error {
	err := m.ms.ChangeSongText(ctx, group, newGroup, name, newName)
	if err != nil {
		if errors.Is(err, ErrConflict) {
			return ErrConflict
		}
		if errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
func (m *ModelMusic) AddSong(ctx context.Context, fsc structs.FullMusicEntry) error {
	err := m.ms.AddSong(ctx, fsc)
	if err != nil {
		if errors.Is(err, ErrConflict) {
			return ErrConflict
		}
		return err
	}

	return nil
}
