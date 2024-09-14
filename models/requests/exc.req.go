package requests

type ExchangeRequest struct {
    ComicIDFrom string `json:"comic_id_from"`
    ComicIDTo   string `json:"comic_id_to"`
    UserIDTo    string `json:"user_id_to"`
}

