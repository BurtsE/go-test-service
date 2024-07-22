package config

type Сonfig struct {
	Postgres `json:"Postgres,omitempty"`
	Service  `json:"Service,omitempty"`
}

type Service struct {
	Port int `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	User     string `env:"PG_USER,notEmpty"`
	Password string `env:"PG_PASSWORD,notEmpty"`
}
