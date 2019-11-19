package templater

import (
	"talk/myAGo/code/templater/test"
)

//START OMIT
//go:generate templater
var _ = test.TemplateA{
	Name:  "Ahoj",
	Value: 42,
}

//END OMIT
