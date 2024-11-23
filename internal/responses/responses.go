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
		Group       string `json:"group" example:"Muse" swagger:"description:Band or artist name"`
		Song        string `json:"song" example:"Supermassive Black Hol" swagger:"description:Song title"`
		ReleaseDate string `json:"releaseDate" validate:"required,min=0" example:"16.07.2006" swagger:"description:Release date"`
		Text        string `json:"text" validate:"required,min=0" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight" swagger:"description:Song lyrics"`
		Link        string `json:"link" validate:"required,min=0" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw" swagger:"description:Link to the song"`
	}

	SongsPaginationResponse struct {
		Songs      []SongInfoResponse `json:"songs" `
		Page       int                `json:"page" example:"1" swagger:"description:Number page"`
		Limit      int                `json:"limit" example:"10" swagger:"description:Limit"`
		Total      int                `json:"total" example:"31" swagger:"description:Count of records found "`
		TotalPages int                `json:"total_pages" example:"4" swagger:"Number of records pages"`
	}
	UpdateResponse struct {
		Message string `json:"message" example:"Update succeeded" swagger:"description:Update succeeded"`
	}
	DeleteResponse struct {
		Message string `json:"message" example:"Delete succeeded" swagger:"description:Delete succeeded"`
	}
)

//Струкутры для ошибок

type (
	ErrorResponse400 struct {
		Code    int    `json:"code" example:"400" `
		Message string `json:"message" example:"Bad Request - Invalid input data"`
	}

	ErrorResponse422 struct {
		Code    int    `json:"code" example:"422" `
		Message string `json:"message" example:"Unprocessable Entity - Validation failed"`
	}

	ErrorResponse404 struct {
		Code    int    `json:"code" example:"404" `
		Message string `json:"message" example:"Not Found - Song not found"`
	}

	ErrorResponse500 struct {
		Code    int    `json:"code" example:"500" `
		Message string `json:"message" example:"Internal Server Error"`
	}
)
