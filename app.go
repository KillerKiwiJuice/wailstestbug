package main

import (
	"context"
	"fmt"
	// "log"
	//"runtime"
	// "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx            context.Context
	test           string
	numClientsPipe chan int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.test = "hello"
	a.numClientsPipe = make(chan int)
	// go handleEvents(a)
	go AllDownhillFromHere(a)
}

// func handleEvents(app *App) {
// 	for {
// 		select {
// 		case message := <-app.numClientsPipe:
// 			log.Println("new client len from app!", message)
// 			runtime.EventsEmit(app.ctx, "onNumClientsChange", message)
// 		}
// 	}
// }

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!!", name)
}

func (a *App) GetTest() string {
	return fmt.Sprintf(a.test)
}
