package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/browser"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (b *App) Authz(service string) string {
	target, err := Authorize(service)
	if err != nil {
		log.Fatalf("Authorize error: %v", err)
	}
	browser.OpenURL(target)

	return fmt.Sprintf("You have been authorized to: %s", service)
}
