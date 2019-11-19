//+build wireinject

package wire

import (
	service "talk/myAGo/code/serviceDep"

	"github.com/google/wire"
)

//START OMIT
//go:generate wire
func CreateApp(config service.Config) (*App, error) {
	wire.Build(
		New,
		service.New,
	)
	return &App{}, nil
}

//END OMIT
