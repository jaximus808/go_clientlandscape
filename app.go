package main

import (
	"context"
	"fmt"
	"palantir_client_landscape/internal/foundry/api"
)

// App struct
type App struct {
	ctx        context.Context
	foundryApi *api.FoundryAPI
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		foundryApi: api.CreateFoundryAPI("/v2/ontologies/ontology-ef92f6b4-0076-41a3-a5c2-780c41133c87/actions/create-leadership-hierachy/apply"),
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
