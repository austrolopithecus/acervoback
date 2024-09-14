package requests

// Estrutura para a requisição de adicionar uma review
type ReviewRequest struct {
	ComicID string `json:"comic_id"`
	UserID  string `json:"user_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

