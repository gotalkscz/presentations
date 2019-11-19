package templater

import (
	"talk/myAGo/code/templater/testFunc"
)

//START OMIT
//go:generate templater
var _ = testFunc.TemplateA{
	Name:  "Ahoj",
	Value: 42,
}

//END OMIT
