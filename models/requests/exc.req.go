package requests

type ExchangeRequest struct {
	ComicIDFrom string `json:"comic_id_from"`
	ComicIDTo   string `json:"comic_id_to"`
	UserIDFrom  string `json:"user_id_from"`
	UserIDTo    string `json:"user_id_to"`
}

