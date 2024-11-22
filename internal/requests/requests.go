package requests

type (
	// @Description Структура запроса для создания новой песни
	SongRequest struct {
		Group string `json:"group" validate:"required,min=0" example:"Muse" swagger:"description:Название группы"`
		Song  string `json:"song" validate:"required,min=0" example:"Supermassive Black Hol" swagger:"description:Название песни"`
	}
	UpdateRequest struct {
		Group       string `json:"group" validate:"required,min=0"`
		Song        string `json:"song" validate:"required,min=0"`
		ReleaseDate string `json:"releaseDate" validate:"required,min=0" `
		Text        string `json:"text" validate:"required,min=0" `
		Link        string `json:"link" validate:"required,min=0" `
	}
)
