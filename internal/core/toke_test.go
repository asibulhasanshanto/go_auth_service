package core

import (
	"testing"
	"time"

	"github.com/asibulhasanshanto/go_api/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// MockConfig returns a mock configuration for testing
func MockConfig() *config.Config {
	return &config.Config{
		App: config.AppConfig{
			AccessTokenSecret:    []byte("access_secret"),
			AccessTokenDuration:  1, // 1 hour
			RefreshTokenSecret:   []byte("refresh_secret"),
			RefreshTokenDuration: 24 * 7, // 1 week
		},
	}
}

func TestGenerateToken(t *testing.T) {
	// Initialize logger
	logger, _ := zap.NewDevelopment()

	// Create a new Token instance with mock config and logger
	token := NewToken(MockConfig(), logger)

	// Define a payload for the token
	payload := map[string]interface{}{
		"user_id": 123,
		"email":   "test@example.com",
	}

	// Generate tokens
	tokens, err := token.GenerateToken(payload)
	logger.Info("tokens", zap.Any("tokens", tokens))
	assert.NoError(t, err, "GenerateToken should not return an error")
	assert.Equal(t, 2, len(tokens), "GenerateToken should return two tokens (access and refresh)")

	// Validate the access token
	accessToken := tokens[0]
	parsedAccessToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return MockConfig().App.AccessTokenSecret, nil
	})
	assert.NoError(t, err, "Access token should be valid")
	assert.True(t, parsedAccessToken.Valid, "Access token should be valid")

	// Validate the refresh token
	refreshToken := tokens[1]
	parsedRefreshToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return MockConfig().App.RefreshTokenSecret, nil
	})
	assert.NoError(t, err, "Refresh token should be valid")
	assert.True(t, parsedRefreshToken.Valid, "Refresh token should be valid")

	// Validate the payload in the access token
	accessTokenClaims, ok := parsedAccessToken.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Access token claims should be of type jwt.MapClaims")
	assert.Equal(t, float64(123), accessTokenClaims["data"].(map[string]interface{})["user_id"], "User ID in access token should match payload")
	assert.Equal(t, "test@example.com", accessTokenClaims["data"].(map[string]interface{})["email"], "Email in access token should match payload")

	// Validate the payload in the refresh token
	refreshTokenClaims, ok := parsedRefreshToken.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Refresh token claims should be of type jwt.MapClaims")
	assert.Equal(t, float64(123), refreshTokenClaims["data"].(map[string]interface{})["user_id"], "User ID in refresh token should match payload")
	assert.Equal(t, "test@example.com", refreshTokenClaims["data"].(map[string]interface{})["email"], "Email in refresh token should match payload")

	// Validate the expiration time of the access token
	accessTokenExp := int64(accessTokenClaims["exp"].(float64))
	assert.True(t, accessTokenExp > time.Now().Unix(), "Access token should not be expired")

	// Validate the expiration time of the refresh token
	refreshTokenExp := int64(refreshTokenClaims["exp"].(float64))
	assert.True(t, refreshTokenExp > time.Now().Unix(), "Refresh token should not be expired")



	// testing the ValidateToken function
	//===================================================================================================
	// Test both access and refresh token validation
	validatedAccessPayload, err := token.ValidateToken(accessToken, "access")
	assert.NoError(t, err, "ValidateToken should not return an error for access token")
	assert.Equal(t, float64(123), validatedAccessPayload["user_id"], "User ID in validated access payload should match")
	assert.Equal(t, "test@example.com", validatedAccessPayload["email"], "Email in validated access payload should match")

	// Test refresh token validation
	validatedRefreshPayload, err := token.ValidateToken(refreshToken, "refresh")
	assert.NoError(t, err, "ValidateToken should not return an error for refresh token")
	assert.Equal(t, float64(123), validatedRefreshPayload["user_id"], "User ID in validated refresh payload should match")
	assert.Equal(t, "test@example.com", validatedRefreshPayload["email"], "Email in validated refresh payload should match")

	// Test with invalid token type
	invalidPayload, err := token.ValidateToken(accessToken, "invalid")
	assert.Empty(t, invalidPayload, "Invalid token type should return empty payload")
	assert.Nil(t, err, "Invalid token type should not return an error")

	// Test with tampered token
	tamperedToken := accessToken + "tampered"
	_, err = token.ValidateToken(tamperedToken, "access")
	assert.Error(t, err, "Tampered token should return an error")

	// Test with wrong token secret (using access token with refresh secret)
	_, err = jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return MockConfig().App.RefreshTokenSecret, nil
	})
	assert.Error(t, err, "Using access token with refresh secret should fail")
}