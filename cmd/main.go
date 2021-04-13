package main

import (
	"fmt"

	devduck "github.com/antklim/dev-duck"
)

func main() {
	fmt.Println("dev-duck")

	router := devduck.Router()
	devduck.Serve(router)
}
