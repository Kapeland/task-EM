package models

type ModelMusic struct {
	ms MusicStorager
}

func NewModelMusic(ms MusicStorager) ModelMusic {
	return ModelMusic{ms}
}
