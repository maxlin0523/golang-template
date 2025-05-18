package main

import (
	"go-template/internal/app"
	_ "go-template/docs"  // 這裡要改成你的 module path
	"go.uber.org/fx"
)


// @title           go template api
// @version         1.0
func main() {
	fx.New(app.Startup).Run()
}
