package responses

type (
	SongResponse struct {
		Song   string   `json:"song"`
		Verses []string `json:"verses"`
	}
	/// @Description Структура ответа при возникновении ошибки
	ErrorResponse struct {
		Code    int    `json:"code" `
		Message string `json:"message"`
	}
	// @Description Response structure containing song information
	SongInfoResponse struct {
		Group       string `json:"group" example:"Beatles" swagger:"description:Band or artist name"`
		Song        string `json:"song" example:"Yesterday" swagger:"description:Song title"`
		ReleaseDate string `json:"releaseDate" validate:"required,min=0" example:"1965-08-06" swagger:"description:Release date"`
		Text        string `json:"text" validate:"required,min=0" example:"Yesterday, all my troubles seemed so far away..." swagger:"description:Song lyrics"`
		Link        string `json:"link" validate:"required,min=0" example:"https://example.com/song" swagger:"description:Link to the song"`
	}
)
