package main

import (
	"fmt"

	pixamgojson "github.com/pixambi/pixam-go-json"
)

func main() {
	var tools pixamgojson.Tools

	s := tools.RandomString(10)
	fmt.Println("Random String:", s)
}
