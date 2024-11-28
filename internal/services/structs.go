package services

// GetSongInfoResp remote endpoint /info
type GetSongInfoResp struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddSongReq struct {
	Group string `json:"group"`
	Name  string `json:"song"`
}

type ChangeSongReq struct {
	Group    string `json:"group"`
	Name     string `json:"song"`
	NewGroup string `json:"new_group"`
	NewName  string `json:"new_song"`
}

type DeleteSongReq struct {
	Group string `json:"group"`
	Name  string `json:"song"`
}

type GetSongTextReq struct {
	Group string `json:"group"`
	Name  string `json:"song"`
}

type GetSongTextResp struct {
	Name string `json:"song"`
	Text string `json:"text"`
}

type FullSongContent struct {
	Group   string `json:"group"`
	Name    string `json:"song"`
	Text    string `json:"text"`
	Release string `json:"release"`
	Link    string `json:"link"`
}

type GetLibraryContentResp struct {
	Library []FullSongContent `json:"songs"`
}
