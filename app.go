package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// App struct
type App struct {
	ctx  context.Context
	mode string
}

// NewApp creates a new App application struct
func NewApp(mode string) *App {
	return &App{
		mode: mode,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Call this method to get current starting mode
func (a *App) GetMode() string {
	return a.mode
}

// TODO: Call this method to run another window
func (a *App) OpenSelf(mode string) error {
	path, _ := os.Executable()

	var cmd string
	var args []string
	cmd = path
	args = []string{"-mode", mode}
	return exec.Command(cmd, args...).Start()
}
