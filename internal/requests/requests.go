package requests

type (
	// @Description Структура запроса для создания новой песни
	SongRequest struct {
		Group string `json:"group" validate:"required,min=0" example:"Muse" swagger:"description:Название группы"`
		Song  string `json:"song" validate:"required,min=0" example:"Supermassive Black Hol" swagger:"description:Название песни"`
	}
	UpdateRequest struct {
		Group       string `json:"group" validate:"required,min=0" example:"Eminem"`
		Song        string `json:"song" validate:"required,min=0" example:"SOng"`
		ReleaseDate string `json:"releaseDate" validate:"required,min=0" example:"00.00.00"`
		Text        string `json:"text" validate:"required,min=0" example:"LaLala"`
		Link        string `json:"link" validate:"required,min=0" example:"http://example.com"`
	}
)
