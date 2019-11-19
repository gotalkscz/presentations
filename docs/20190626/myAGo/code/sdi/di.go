package sdi

import (
	"talk/myAGo/code/lib/s"
	"talk/myAGo/code/service"
	app "talk/myAGo/code/wireFormaterTrabl"
)

//START OMIT
//go:generate templater
var _ = func() *s.Di {
	return s.NewDi("app").
		Add(
			app.New,
			app.DefaultFormater,
			s.Scope(
				s.Config(service.NewConfig, "helloA"),
				s.Service(service.New),
			),
			s.Scope(
				s.Config(service.NewConfig, "helloB"),
				s.Service(service.New),
			),
		)
}

//END OMIT
