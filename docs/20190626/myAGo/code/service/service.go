package service

type Formater interface {
	Format(string) string
}

type Hello struct {
	formater Formater
	config   Config
}

func New(config Config, formater Formater) *Hello { // HL1
	return &Hello{
		config:   config,
		formater: formater,
	}
}

func (h *Hello) Say() string {
	return h.formater.Format(h.config.Name)
}
