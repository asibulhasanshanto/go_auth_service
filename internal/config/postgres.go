package config

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB   string
	MaxIdleConns int
	MaxOpenConns int
	ConnMaxLifetime int
}

// GetDSN returns the Data Source Name (DSN) for database connection
func (dc *PostgresConfig) GetDSN() string {
	return "host=" + dc.Host + " user=" + dc.User + " password=" + dc.Password + " dbname=" + dc.DB + " port=" + dc.Port  + " sslmode=disable"
}