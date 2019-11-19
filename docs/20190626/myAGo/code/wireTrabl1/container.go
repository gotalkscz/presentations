//+build wireinject

package wireTrabl1

import (
	"fmt"

	"talk/myAGo/code/service"

	"github.com/google/wire"
)

//START OMIT
//go:generate wire
func CreateApp(_ service.Config) (*App, error) {
	wire.Build(
		New,
		service.New,
		NewServiceA,
		NewServiceB,
		DefaultFormater,
	)
	return &App{}, nil
}

//END OMIT

type FmtFormater string // HLnew

func (f FmtFormater) Format(in string) string {
	return fmt.Sprintf(string(f), in)
} // HLnew

func DefaultFormater() service.Formater { // HLnew
	return FmtFormater("jméno: %s")
}
