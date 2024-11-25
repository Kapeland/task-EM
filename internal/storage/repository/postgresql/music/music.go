package tests

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/storage/db"
)

type Repo struct {
	db db.DBops
}

func New(db db.DBops) *Repo {
	return &Repo{db: db}
}

func (m *Repo) GetMusicText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
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
