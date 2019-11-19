package main

import (
	"flag"
	"fmt"
)

func main() {
	value := "defaultní hodnota"
	flag.StringVar(&value, "value", value, "string hodnota")
	flag.Parse()
	fmt.Println(value)
}
