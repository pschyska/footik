package controllers

import "github.com/robfig/revel"

type Application struct {
	*rev.Controller
}

func (c Application) Index() rev.Result {
  greeting := "Hello World!"
	return c.Render(greeting)
}

func (c Application) Hello(myName string) rev.Result {
  return c.Render(myName)
}
