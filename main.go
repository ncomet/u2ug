package main

import (
	"fmt"
	"os"

	"github.com/ncomet/u2ug/ubi"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expecting UUID argument, usage:")
		fmt.Println("u2ug a496d62f-b0bc-4c58-96b9-ed838127d724")
		fmt.Println("exiting")
		os.Exit(1)
	}
	game, err := ubi.Game(os.Args[1])
	if err != nil {
		fmt.Printf("error: %s, exiting\n", err)
		os.Exit(1)
	}
	fmt.Println(game)
	os.Exit(0)
}
