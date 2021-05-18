package main

import (
	"fmt"
)

// TODO: add https://github.com/oklog/run

func main() {
	fmt.Println("dev-duck")

	r := Router()
	Serve(r)
}
