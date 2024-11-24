package responses

type UserLoginResponse struct {
	CommonResponse
	Token string `json:"token"`
}

type UserMeResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
