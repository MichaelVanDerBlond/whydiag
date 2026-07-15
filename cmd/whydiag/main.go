package main

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/whydiag/internal/app"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
