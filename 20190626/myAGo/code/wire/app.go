package wire

import (
	"fmt"
	service "talk/myAGo/code/serviceDep"
)

//START OMIT
type App struct{ hello *service.Hello }

func New(hello *service.Hello) *App {
	return &App{
		hello: hello,
	}
}

func (a *App) Run() {
	fmt.Println(a.hello.Say())
}

//END OMIT
