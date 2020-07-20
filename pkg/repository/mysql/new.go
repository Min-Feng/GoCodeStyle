package mysql

type Config struct {
	User     string `configs:"user"`
	Password string `configs:"password"`
	Host     string `configs:"host"`
	Port     string `configs:"port"`
}

func New(cfg Config) {
}
