package testFunc

import (
	"fmt"
)

type TemplateA struct {
	Name  string
	Value int
}

func (a TemplateA) Full() string {
	return fmt.Sprintf("%s (%d)", a.Name, a.Value)
}
