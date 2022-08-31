//go:build wasm

package main

import (
	"syscall/js"

	"github.com/ncomet/u2ug/ubi"
)

func ubiGame(this js.Value, p []js.Value) any {
	game, _ := ubi.Game(p[0].String())
	return game
}

func ubiCharacter(this js.Value, p []js.Value) any {
	char, _ := ubi.Character(p[0].String())
	return char
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("ubiGame", js.FuncOf(ubiGame))
	js.Global().Set("ubiCharacter", js.FuncOf(ubiCharacter))
	<-c
}
