package main

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/whydiag/internal/app"
	"github.com/MichaelVanDerBlond/whydiag/internal/version"
)

func usage() {
	fmt.Println("WhyDiag")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  whydiag check")
	fmt.Println("  whydiag version")
}

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {

	case "version":
		fmt.Printf("%s %s\n", version.Name, version.Version)

	case "check":
		a := app.New()

		if err := a.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		usage()
		os.Exit(1)
	}
}
