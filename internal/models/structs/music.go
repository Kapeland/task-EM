package structs

type LibraryEntry struct {
	ID int `db:"song_group"`
}

type MusicEntry struct {
	Name string `db:"song"`
	Text string `db:"song_text"`
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
