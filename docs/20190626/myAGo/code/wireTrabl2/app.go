package wireFormaterTrabl

import (
	"fmt"
	"talk/myAGo/code/service"
)

func New(a *ServiceA, b *ServiceB) *App {
	return &App{
		a: a,
		b: b,
	}
}

//START OMIT
type ServiceAHello service.Hello
type ServiceBHello service.Hello
type ConfigAHello service.Config
type ConfigBHello service.Config

func NewServiceHelloA(config ConfigAHello, formater service.Formater) *ServiceAHello {
	return (*ServiceAHello)(service.New(service.Config(config), formater))
}

func NewServiceHelloB(config ConfigBHello, formater service.Formater) *ServiceBHello {
	return (*ServiceBHello)(service.New(service.Config(config), formater))
}

//END OMIT

type App struct {
	a *ServiceA
	b *ServiceB
}

type ServiceA struct {
	say *service.Hello
}

func NewServiceA(say *ServiceAHello) *ServiceA {
	return &ServiceA{
		say: (*service.Hello)(say),
	}
}

type ServiceB struct {
	say *service.Hello
}

func NewServiceB(say *ServiceBHello) *ServiceB {
	return &ServiceB{
		say: (*service.Hello)(say),
	}
}

func (a *ServiceA) Run() {
	fmt.Println(a.say.Say())
}

func (a *ServiceB) Run() {
	fmt.Println(a.say.Say())
}

type FmtFormater string // HLnew

func (f FmtFormater) Format(in string) string {
	return fmt.Sprintf(string(f), in)
}

func DefaultFormater() service.Formater {
	return FmtFormater("jméno: %s")
}
