package app

import (
	"fmt"

	"github.com/MichaelVanDerBlond/whydiag/internal/checks/linux"
	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/version"
)

type App struct {
	runner *core.Runner
}

func New() *App {
	r := core.NewRunner()

	r.Register(linux.New())

	return &App{
		runner: r,
	}
}

func (a *App) Run() error {
	fmt.Printf("%s %s\n\n", version.Name, version.Version)

	a.runner.Run()

	return nil
}
