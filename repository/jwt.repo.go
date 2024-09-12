package repository

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"time"
)

var secretKey = []byte("your-256-bit-secret")

type JWTRepo interface {
	GenerateToken(id string) (string, error)
	VerifyToken(token string) (string, error)
}

type JwtRepoImpl struct {
}

func (j *JwtRepoImpl) GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(secretKey)
}

func (j *JwtRepoImpl) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Valida o token com o método de assinatura HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error().Msg("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		log.Err(err).Msg("Erro ao verificar token")
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Debug().Str("id", claims["id"].(string)).Msg("Token valido")
		return claims["id"].(string), nil
	} else {
		log.Error().Msg("token invalido")
		return "", fmt.Errorf("invalid token")
	}
}
