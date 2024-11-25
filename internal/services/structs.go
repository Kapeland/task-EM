package services

type GetAllTestsReq struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type GetAllTestsResp struct {
	Tests []GetAllTestsRespTest `json:"tests"`
}

type GetAllTestsRespTest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DiffLevel   int    `json:"diff_level"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Picture     string `json:"picture"`
}

type GetFullTestResp struct {
	Id          int                       `json:"id"`
	Name        string                    `json:"name"`
	DiffLevel   int                       `json:"diff_level"`
	Description string                    `json:"description"`
	Category    string                    `json:"category"`
	Questions   []GetFullTestRespQuestion `json:"questions"`
}

type GetFullTestRespQuestion struct {
	Id       int                     `json:"id"`
	Question string                  `json:"question"`
	IsSong   bool                    `json:"isSong"`
	Song     string                  `json:"song"`
	Answers  []GetFullTestRespAnswer `json:"answers"`
}

type GetFullTestRespAnswer struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

type GetUserScoreReq struct {
	Answers []GetUserScoreReqAnswer `json:"user_answers"`
}

type GetUserScoreReqAnswer struct {
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}

type GetUserScoreResp struct {
	UserScore int `json:"user_score"`
	Total     int `json:"total"`
}

type GetRatingResp struct {
	Rating []GetRatingRespUnit `json:"rating"`
}

type GetRatingRespUnit struct {
	Login string `json:"login"`
	Place int    `json:"place"`
	Score int    `json:"score"`
}
