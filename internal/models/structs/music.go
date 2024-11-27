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

type TestFull struct {
	ID          int
	Name        string
	Description string
	Category    string
	DiffLevel   int
	Questions   []TestQuestionFull
}

type TestQuestionFull struct {
	ID       int
	TestID   int
	Question string
	IsSong   bool
	Song     []byte
	Answers  []QuestionAnswer
}

type QuestionAnswer struct {
	ID         int    `db:"id"`
	QuestionID int    `db:"question_id"`
	Answer     string `db:"answer"`
	IsCorrect  bool   `db:"is_correct"`
}
