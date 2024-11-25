package storage

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models/structs"
)

type MusicRepo interface {
	GetAllMusic(ctx context.Context, id int) (structs.TestFull, error)
	GetMusicText(ctx context.Context, id int) (structs.TestFull, error)
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

func (m *MusicStorage) GetSongText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}

func (t *MusicStorage) GetAllMusic(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}

func (t *MusicStorage) DeleteSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (t *MusicStorage) AddSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (t *MusicStorage) ChangeSongText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
