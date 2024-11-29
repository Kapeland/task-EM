package structs

import "time"

type MusicEntry struct {
	Name string `db:"song"`
	Text string `db:"song_text"`
}

type FullMusicEntry struct {
	Group   string    `db:"song_group"`
	Name    string    `db:"song"`
	Text    string    `db:"song_text"`
	Release time.Time `db:"release_date"`
	Link    string    `db:"link"`
}
