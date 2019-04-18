package main

import (
	"flag"
	"fmt"
)

func main() {

	var subject = flag.String("n", "bob", "the user")
	flag.Parse()

	fmt.Printf("Hello %s from the 0.0.1 branch ! \n", *subject)

	fmt.Print("This appeared in the 0.0.2 tag ! \n ")

	fmt.Printf("Some minor mods \n")

}
