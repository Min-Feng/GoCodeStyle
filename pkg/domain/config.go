package domain

type Config struct {
	MySQLUser     string
	MySQLPassword string
	MailAddress   string
}

type ConfigStore interface {
	Find() (Config, error)
}
