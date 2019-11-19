//+build wireinject

package wireFormater

import (
	"fmt"

	"talk/myAGo/code/service"
)

//START OMIT

type FmtFormater string // HLnew

func (f FmtFormater) Format(in string) string {
	return fmt.Sprintf(string(f), in)
}

func DefaultFormater() service.Formater { // HLnew
	return FmtFormater("jméno: %s")
}

//go:generate wire
func CreateApp(config service.Config) (*App, error) {
	wire.Build(
		New,
		service.New,
		DefaultFormater, // HLnew
	)
	return &App{}, nil
}

//END OMIT
