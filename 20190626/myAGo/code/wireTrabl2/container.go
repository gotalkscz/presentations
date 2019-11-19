//+build wireinject

package wireFormaterTrabl

import (
	"github.com/google/wire"
)

//START OMIT
//go:generate wire
func CreateApp(a ConfigAHello, b ConfigBHello) (*App, error) {
	wire.Build(
		New,
		NewServiceA,
		NewServiceB,
		NewServiceHelloA,
		NewServiceHelloB,
		//END OMIT
		DefaultFormater, // HLnew
	)
	return &App{}, nil
}

//END OMIT
