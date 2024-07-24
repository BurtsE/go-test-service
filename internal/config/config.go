package config

type Config struct {
	Postgres `json:"postgres,omitempty"`
	Service  `json:"service,omitempty"`
}

type Service struct {
	Port string `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	User     string `env:"PG_USER,notEmpty"`
	Password string `env:"PG_PASSWORD,notEmpty"`
}
