package main

import (
	"fmt"

	"github.com/teris-io/shortid"
)

func main() {
	id := shortid.GetDefault(). MustGenerate()
	fmt.Println(id)
}