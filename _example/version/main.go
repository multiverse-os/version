package main

import (
	"fmt"

	version "../../../version"
)

func main() {
	fmt.Println("Simple Version Usage")
	fmt.Println("====================")

	fmt.Println("Setting up a version object for 4.2.0...")
	v := version.Version{Major: 4, Minor: 2, Patch: 0}

	fmt.Println("Now using the String() function to print it: ", v.String())

}
