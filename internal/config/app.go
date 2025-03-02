package config

type AppConfig struct {
	Port string
	AccessTokenSecret []byte
	RefreshTokenSecret []byte
	AccessTokenDuration int
	RefreshTokenDuration int
}