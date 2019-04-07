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

func (c App) Param(id int, str string) revel.Result {

	/*
	   Two other ways to pull the parameter out of the query string

	   var id string = c.Params.Get("id")

	   Or

	   var id int
	   c.Params.Bind(&id, "id")
	*/

	/*
		This doesn't quite work, the cookie gets set and unset at slightly wrong time
	*/
	c.Validation.Required(id).Message("ID is required!")
	c.Validation.Required(str).Message("String parameter missing from URL!")
	c.Validation.MinSize(str, 3).Message("String is not long enough!")

	if c.Validation.HasErrors() {
		// Sets the flash parameter `error` which will be sent by a flash cookie
		c.Flash.Error("Something went wrong!")
		// Keep the validation error from above by setting a flash cookie
		//		c.Validation.Keep()
		// Copies all given parameters (URL, Form, Multipart) to the flash cookie
		//		c.FlashParams()
	}
	return c.Render(id, str)
}

func (c App) NoSlash() revel.Result {
	return c.Render()
}

func (c App) Slash() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	// Explanation of Flash
	// https://revel.github.io/manual/sessionflash.html#flash

	c.Validation.Required(myName).Message("Name is required!")
	c.Validation.MinSize(myName, 3).Message("Name is not long enough!")

	if c.Validation.HasErrors() {
		// Sets the flash parameter `error` which will be sent by a flash cookie
		c.Flash.Error("Something went wrong!")
		// Keep the validation error from above by setting a flash cookie
		c.Validation.Keep()
		// Copies all given parameters (URL, Form, Multipart) to the flash cookie
		c.FlashParams()
		return c.Redirect(App.Index)
	}

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
