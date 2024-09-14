package models

type ExchangeStatus string

const (
    Pending  ExchangeStatus = "pending"
    Accepted ExchangeStatus = "accepted"
    Declined ExchangeStatus = "declined"
    Completed ExchangeStatus = "completed"
)

type Exchange struct {
    ID            string         `json:"id" gorm:"primaryKey"`
    ComicIDFrom   string         `json:"comic_id_from"`   // Quadrinho que o usuário está oferecendo
    ComicIDTo     string         `json:"comic_id_to"`     // Quadrinho que o outro usuário está oferecendo
    UserIDFrom    string         `json:"user_id_from"`    // Usuário que está iniciando a troca
    UserIDTo      string         `json:"user_id_to"`      // Usuário que está recebendo a oferta
    Status        ExchangeStatus `json:"status"`          // Status da troca (pendente, aceita, etc.)
}


