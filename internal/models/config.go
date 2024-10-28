package models

type Config struct {
	Mode string
	DB   ConfigDB
	HTTP ConfigHTTP
	Auth ConfigAuth
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
