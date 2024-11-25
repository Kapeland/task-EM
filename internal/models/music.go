package models

import (
	"context"
	"github.com/Kapeland/task-EM/internal/models/structs"
)

type MusicStorager interface {
	GetAllMusic(ctx context.Context, id int) (structs.TestFull, error)
	GetSongText(ctx context.Context, id int) (structs.TestFull, error)
	DeleteSong(ctx context.Context, id int) (structs.TestFull, error)
	AddSong(ctx context.Context, id int) (structs.TestFull, error)
	ChangeSongText(ctx context.Context, id int) (structs.TestFull, error)
}

const defaultLimit = 100
const defaultOffset = 0

const defaultRatingLimit = 200
const defaultRatingOffset = 0

const TotalCategory = ""

func (m *ModelMusic) GetLibInfo(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *ModelMusic) GetSongText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *ModelMusic) DeleteSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *ModelMusic) ChangeSongText(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}
func (m *ModelMusic) AddSong(ctx context.Context, id int) (structs.TestFull, error) {
	return structs.TestFull{}, nil
}

//func (m *ModelMusic) GetMusic(ctx context.Context, category string, limit int, offset int) ([]structs.TestSimple, error) {
//	if limit == 0 {
//		limit = defaultLimit
//	}
//	if offset == 0 {
//		offset = defaultOffset
//	}
//
//	if category == "" {
//		category = TotalCategory
//	}
//
//	tests, err := m.ms.GetMusic(ctx, category, limit, offset)
//	if err != nil {
//		return nil, err
//	}
//
//	return tests, nil
//}
