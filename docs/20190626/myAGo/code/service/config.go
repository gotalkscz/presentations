package service

type Config struct {
	Name string
}

func NewConfig() Config {
	return Config{
		Name: "Honza",
	}
}
