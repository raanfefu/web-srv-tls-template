package server

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	c "github.com/raanfefu/web-srv-tls-template/common"
)

type WebServer interface {
	CheckServerParams()
	InitServer()
	AddEndpoint(path string, handler func(http.ResponseWriter, *http.Request), methods ...string)
	StartServer()
}

type Impl struct {
	Params *ServerParams
	Server *http.Server
	Router *mux.Router
}

func NewServer() WebServer {

	return &Impl{
		Params: &ServerParams{},
	}
}

func (s *Impl) CheckServerParams() {

	var mode string
	flag.StringVar(&mode, "mode", "https", "mode value is https / http")
	flag.StringVar(&s.Params.Key, "key", "", "absoluted path file key pem")
	flag.StringVar(&s.Params.Crt, "crt", "", "absoluted path file cert pem")
	flag.Int64Var(&s.Params.Port, "port", 443, "port using listing service")
	flag.Parse()

	if c.StringIsRequiered(&mode) != nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	vmode, err := ParseMode(&mode)
	if err != nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	s.Params.Mode = *vmode

	if c.StringIsRequiered(&s.Params.Key) != nil && s.Params.Mode == TLS {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if c.StringIsRequiered(&s.Params.Crt) != nil && s.Params.Mode == TLS {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if c.Int64IsRequiered(s.Params.Port) != nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	log.Printf("Reading parameters... Done ✓")
}

func (s *Impl) InitServer() {

	s.Router = mux.NewRouter()
	switch mode := s.Params.Mode; mode {
	case TLS:
		certificate := s.loadCertificate()
		s.Server = &http.Server{
			Handler:   s.Router,
			Addr:      fmt.Sprintf(":%v", s.Params.Port),
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{certificate}},
		}
	case HTTP:
		s.Server = &http.Server{
			Handler: s.Router,
			Addr:    fmt.Sprintf(":%v", s.Params.Port),
		}
	}
	log.Println("Initializing server... Done ✓")
}

func (s *Impl) AddEndpoint(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {

	s.Router.HandleFunc(path, handler).Methods(methods...)
	log.Printf("Adding endpoint: Resoure name: %s Method: %s ... Done ✓", path, methods)
}

func (s *Impl) StartServer() {
	if s.Params.Mode == TLS {
		go func() {
			if err := s.Server.ListenAndServeTLS("", ""); err != nil {
				fmt.Printf("Failed to listen and serve webhook server: %v\n", err)
			}
		}()
	} else {
		go func() {
			if err := s.Server.ListenAndServe(); err != nil {
				fmt.Printf("Failed to listen and serve webhook server: %v\n", err)
			}
		}()
	}
	log.Printf("Starting web server 📡... Done ✓")
	log.Printf("Listen on %s://0.0.0.0:%v/", s.Params.Mode.String(), s.Params.Port)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	fmt.Println("Got shutdown signal, shutting down webhook server gracefully...\n")
	s.Server.Shutdown(context.Background())
}

func (s *Impl) loadCertificate() tls.Certificate {
	certs, err := tls.LoadX509KeyPair(s.Params.Crt, s.Params.Key)
	if err != nil {
		log.Panicf("Failed to load key pair: %v\n", err)
	}
	log.Printf("Loading certificate TLS ... Done ✓")
	return certs
}
