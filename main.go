package main

import (
	"flag"
	"fmt"
)

func main() {

	var subject = flag.String("n", "bob", "the user")
	flag.Parse()

	fmt.Printf("Hello %s ! \n", *subject)

}
