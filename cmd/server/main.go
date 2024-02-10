package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/raanfefu/web-srv-tls-template/pkg/dummy"
	"github.com/raanfefu/web-srv-tls-template/pkg/server"
)

func show() {
	ofigure := figure.NewColorFigure("web server", "", "green", true)
	ofigure.Print()
	fmt.Printf("\nDeveloper: Rafael Fernández 🐼🐼🐼\n\n")
}

func main() {
	show()
	s := server.NewServer()
	s.CheckServerParams()
	s.InitServer()
	s.AddEndpoint("/a", dummy.DummyHandler, "GET", "POST")
	s.StartServer()
}
