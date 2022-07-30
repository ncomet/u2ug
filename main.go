package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ncomet/u2ug/ubi"
)

var (
	game      = flag.NewFlagSet("game", flag.ExitOnError)
	character = flag.NewFlagSet("character", flag.ExitOnError)
)

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("expected 'game' or 'character' subcommands")
		os.Exit(1)
	}
	var genFunc func(string) (string, error)
	switch os.Args[1] {
	case "game":
		genFunc = ubi.Game
	case "character":
		genFunc = ubi.Character
	}

	if len(os.Args) < 3 {
		fmt.Println("expecting UUID argument, usage:")
		fmt.Println("u2ug game a496d62f-b0bc-4c58-96b9-ed838127d724\nor\nu2ug character a496d62f-b0bc-4c58-96b9-ed838127d724")
		fmt.Println("exiting")
		os.Exit(1)
	}
	gen, err := genFunc(os.Args[2])
	if err != nil {
		fmt.Printf("error: %s, exiting\n", err)
		os.Exit(1)
	}
	fmt.Println(gen)
	os.Exit(0)
}
