package config

type AppConfig struct {
	Port                 string
	AccessTokenSecret    string // Changed from []byte to string
	RefreshTokenSecret   string // Changed from []byte to string
	AccessTokenDuration  int
	RefreshTokenDuration int
}
