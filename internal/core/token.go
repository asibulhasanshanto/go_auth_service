package core

import (
	"time"

	"github.com/asibulhasanshanto/go_api/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Token struct {
	cfg *config.Config
	log *zap.Logger
}

func NewToken(
	cfg *config.Config,
	log *zap.Logger,
	) *Token {
	return &Token{
		cfg: cfg,
		log: log,	
	}
}

func (t *Token) GenerateToken(payload map[string]interface{}) ([]string, error) {
	accessTokenSecret := t.cfg.App.AccessTokenSecret
	accessTokenDuration := t.cfg.App.AccessTokenDuration
	refreshTokenSecret := t.cfg.App.RefreshTokenSecret
	refreshTokenDuration := t.cfg.App.RefreshTokenDuration

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{ 
		"data": payload,
		"exp": time.Now().Add(time.Duration(accessTokenDuration) * time.Hour).Unix(), 
		})

		signedAccessToken, err := accessToken.SignedString(accessTokenSecret)
		if err != nil {
			return []string{}, err
		}

   refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
   jwt.MapClaims{
   "data": payload,
   "exp": time.Now().Add(time.Duration(refreshTokenDuration) * time.Hour).Unix(),
   })

   signedRefreshToken, err := refreshToken.SignedString(refreshTokenSecret)
   if err != nil {
	   return []string{}, err
   }

   return []string{signedAccessToken, signedRefreshToken}, nil
}

func (t *Token) ValidateToken(tokenString string, tokenType string) (map[string]interface{}, error) {
	var secret []byte
	if tokenType == "access" {
		secret = t.cfg.App.AccessTokenSecret
	} else if tokenType == "refresh" {
		secret = t.cfg.App.RefreshTokenSecret
	} else {
		return map[string]interface{}{}, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return map[string]interface{}{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return map[string]interface{}{}, err
	}

	return claims["data"].(map[string]interface{}), nil
}