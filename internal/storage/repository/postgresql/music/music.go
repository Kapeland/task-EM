package tests

import (
	"context"
	"database/sql"
	"github.com/Kapeland/task-EM/internal/models"
	"github.com/Kapeland/task-EM/internal/models/structs"
	"github.com/Kapeland/task-EM/internal/storage/db"
	"github.com/Kapeland/task-EM/internal/storage/repository"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"strings"
)

type Repo struct {
	db db.DBops
}

func New(db db.DBops) *Repo {
	return &Repo{db: db}
}

// GetMusicText directly extracts song name and text from postgres
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
	song.Text = strings.ReplaceAll(song.Text, "\\n", "\n")
	return song, nil
}

// GetAllMusic directly extracts all library from postgres
func (m *Repo) GetAllMusic(ctx context.Context, group string) ([]structs.FullMusicEntry, error) {
	if group == models.GroupFilter {
		group = "%"
	}

	var songs []*structs.FullMusicEntry
	err := m.db.Select(ctx, &songs,
		`SELECT song_group, song, song_text, release_date, link
		FROM library_schema.library
		WHERE song_group like $1;`, group)
	if err != nil {
		return nil, err
	}
	songsOut := make([]structs.FullMusicEntry, len(songs))
	for i, song := range songs {
		songsOut[i] = *song
		songsOut[i].Text = strings.ReplaceAll(song.Text, "\\n", "\n")
	}
	return songsOut, nil
}

// DelSong directly deletes song in postgres
func (m *Repo) DelSong(ctx context.Context, group string, name string) error {
	tag, err := m.db.Exec(ctx,
		`DELETE FROM library_schema.library WHERE song_group = $1 and song = $2;`, group, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return repository.ErrObjectNotFound
		}
		return err
	}
	if tag.RowsAffected() == 0 {
		return repository.ErrObjectNotFound
	}
	return nil
}

// PutSong directly modifies song name and group in postgres
func (m *Repo) PutSong(ctx context.Context, group string, newGroup string, name string, newName string) error {
	tmp := ""
	err := m.db.ExecQueryRow(ctx,
		`UPDATE library_schema.library set 
				song_group = $1, song = $2 WHERE song_group = $3 and song = $4 returning song;`, newGroup, newName, group, name).Scan(&tmp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return repository.ErrObjectNotFound
		}
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			return repository.ErrDuplicateKey
		}
		return err
	}
	return nil
}

// PostSong directly adds song to in postgres
func (m *Repo) PostSong(ctx context.Context, fsc structs.FullMusicEntry) error {
	tmp := ""
	err := m.db.ExecQueryRow(ctx,
		`INSERT INTO library_schema.library(song_group, song, song_text, release_date, link)
				VALUES($1,$2,$3,$4,$5) returning song;`, fsc.Group, fsc.Name, fsc.Text, fsc.Release, fsc.Link).Scan(&tmp)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			return repository.ErrDuplicateKey
		}
		return err
	}

	return nil
}
