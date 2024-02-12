package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/raanfefu/web-srv-tls-template/pkg/cfg"
	"github.com/raanfefu/web-srv-tls-template/pkg/server"
)

func show() {
	ofigure := figure.NewColorFigure("web server", "", "green", true)
	ofigure.Print()
	fmt.Printf("\nDeveloper: Rafael Fernández 🐼🐼🐼\n\n")
}

func main() {
	show()
	cfgs := cfg.NewConfiguration()
	server := server.NewServer()
	cfgs.RegistryService("w", server)
	cfgs.LoadConfiguration()
	server.InitServer()
	server.StartServer()
}
