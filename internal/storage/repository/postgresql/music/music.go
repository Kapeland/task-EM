package tests

import (
	"context"
	"database/sql"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/storage/db"
	"github.com/Kapeland/task-EM/internal/storage/repository"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Repo struct {
	db db.DBops
}

func New(db db.DBops) *Repo {
	return &Repo{db: db}
}

func (m *Repo) GetMusicText(ctx context.Context, group string, name string) (structs.MusicEntry, error) {
	var song structs.MusicEntry

	err := m.db.Get(ctx, &song,
		`Select song, song_text FROM library_schema.library WHERE song_group = $1 and song = $2;`, group, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return structs.MusicEntry{}, repository.ErrObjectNotFound
		}
		return structs.MusicEntry{}, err
	}
	return song, nil
}

func (m *Repo) GetAllMusic(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}

func (m *Repo) DelSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *Repo) PutSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *Repo) PostSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
