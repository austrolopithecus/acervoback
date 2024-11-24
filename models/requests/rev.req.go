package requests

// Estrutura para a requisição de adicionar uma review
type ReviewRequest struct {
	ComicID string `json:"comic_id"`
	UserID  string `json:"user_id"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}
