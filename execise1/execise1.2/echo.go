package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("Index %d and value %s\n", i, v)
	}
}
