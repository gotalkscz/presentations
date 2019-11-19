// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wireFormaterTrabl

// Injectors from container.go:
//START OMIT
func CreateApp(a ConfigAHello, b ConfigBHello) (*App, error) {
	formater := DefaultFormater()
	serviceAHello := NewServiceHelloA(a, formater)
	serviceA := NewServiceA(serviceAHello)
	serviceBHello := NewServiceHelloB(b, formater)
	serviceB := NewServiceB(serviceBHello)
	app := New(serviceA, serviceB)
	return app, nil
}

//END OMIT
