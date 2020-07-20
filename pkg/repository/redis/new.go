package redis

type Config struct {
	User     string   `configs:"user"`
	Password string   `configs:"password"`
	Address  []string `configs:"address"`
}

func New(cfg Config) {
}
