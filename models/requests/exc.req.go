package requests

type ExchangeRequest struct {
	ComicID     string `json:"comic_id"`
	RequesterID string `json:"requester_id"`
	OwnerID     string `json:"owner_id"`
}
