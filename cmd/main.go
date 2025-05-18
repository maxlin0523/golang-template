package main

import (
	"go-template/internal/app"
	"go.uber.org/fx"
)


// @title           go template api
// @version         1.0
func main() {
	fx.New(app.Startup).Run()
}
