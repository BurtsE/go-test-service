package config

type Config struct {
	Postgres `json:"postgres,omitempty"`
	Service  `json:"service,omitempty"`
	Kafka    `json:"kafka,omitempty"`
}

type Service struct {
	Port string `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	Sslmode  string `json:"sslmode"`
	User     string `env:"PG_USER,notEmpty"`
	Password string `env:"PG_PASSWORD,notEmpty"`
}

type Kafka struct {
	Brokers []string `json:"brokers,omitempty"`
	Topic   string   `json:"topic,omitempty"`
}
