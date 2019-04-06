package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}

// Whatever arguments are written here, the framework automatically
// grabs them from the query string and builds the code to route them in
func (c App) Robin(first string, third string) revel.Result {
	// this is overriding the query string parameter called first if it is empty
	if first == "" {
		first = "No first passed"
	}
	// This is making up a new parameter called second
	second := "second"
	// The names of the parameters here end up as named
	// parameters in the view
	// Content written out is automatically HTML encoded
	return c.Render(first, second, third)
}
