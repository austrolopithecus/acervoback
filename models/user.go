package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
type UserClaims struct {
	jwt.MapClaims
	ID string `json:"id"`
}
