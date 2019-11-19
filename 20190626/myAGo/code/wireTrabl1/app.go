package wireTrabl1

import (
	"fmt"
	"talk/myAGo/code/service"
)

type App struct {
	a *ServiceA
	b *ServiceB
}

func New(a *ServiceA, b *ServiceB) *App {
	return &App{
		a: a,
		b: b,
	}
}

//START OMIT
type ServiceA struct {
	say *service.Hello
}

func NewServiceA(say *service.Hello) *ServiceA {
	return &ServiceA{
		say: say,
	}
}

type ServiceB struct {
	say *service.Hello
}

func NewServiceB(say *service.Hello) *ServiceB {
	return &ServiceB{
		say: say,
	}
}

//END OMIT

func (a *ServiceA) Run() {
	fmt.Println(a.say.Say())
}

func (a *ServiceB) Run() {
	fmt.Println(a.say.Say())
}
