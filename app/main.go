package main

import (
	"github.com/brunobotter/feature-flag/main/app"
	"github.com/brunobotter/feature-flag/main/providers"
)

func main() {
	app.NewApplication(providers.List()).Bootstrap()
}
