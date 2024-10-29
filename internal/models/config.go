package models

type AppMode string

const (
	AppModeProduction AppMode = "production"
	AppModeDevelop    AppMode = "develop"
)

type Config struct {
	Mode AppMode
	DB   ConfigDB
	HTTP ConfigHTTP
	Auth ConfigAuth
}

func (c *Config) IsDevelop() bool {
	return c.Mode == AppModeDevelop
}

type ConfigDB struct {
	DSN string
}

type ConfigHTTP struct {
	Port    uint16
	BaseURL string
}

type ConfigAuth struct {
	JWTSecretKey string
}
