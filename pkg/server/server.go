package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/raanfefu/web-srv-tls-template/pkg/cfg"
	"github.com/raanfefu/web-srv-tls-template/pkg/commons"
)

type WebServer interface {
	InitServer()
	AddEndpoint(path string, handler func(http.ResponseWriter, *http.Request), methods ...string)
	StartServer()
}

type ServerParams struct {
	Mode  *commons.ModeType
	Port  uint
	Certs tls.Certificate
}

type webServerImpl struct {
	cfg.DefaultConfiguraionService
	Params *ServerParams
	Server *http.Server
	Router *mux.Router
	mode   string
}

func NewServer() WebServer {
	return &webServerImpl{
		Params: &ServerParams{},
	}
}

func (s *webServerImpl) InitServer() {
	if s.Params.Mode == nil {
		os.Exit(-1)
	}
	mode := *s.Params.Mode
	s.Router = mux.NewRouter()
	switch mode {
	case commons.TLS:
		if s.Params.Port == 0 {
			s.Params.Port = 443
		}
		certificate := s.Params.Certs
		s.Server = &http.Server{
			Handler:   s.Router,
			Addr:      fmt.Sprintf(":%v", s.Params.Port),
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{certificate}},
		}
	case commons.HTTP:
		if s.Params.Port == 0 {
			s.Params.Port = 80
		}
		s.Server = &http.Server{
			Handler: s.Router,
			Addr:    fmt.Sprintf(":%v", s.Params.Port),
		}
	}
	log.Println("Initializing server... Done ✓")
}

func (s *webServerImpl) AddEndpoint(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {

	s.Router.HandleFunc(path, handler).Methods(methods...)
	log.Printf("Adding endpoint: Resoure name: %s Method: %s ... Done ✓", path, methods)
}

func (s *webServerImpl) StartServer() {
	if s.Params.Mode == nil {
		os.Exit(-1)
	}
	mode := *s.Params.Mode
	if mode == commons.TLS {
		go func() {
			if err := s.Server.ListenAndServeTLS("", ""); err != nil {
				fmt.Printf("Failed to listen and serve webhook server: %v\n", err)
				os.Exit(-1)
			}
		}()
	} else {
		go func() {
			if err := s.Server.ListenAndServe(); err != nil {
				fmt.Printf("Failed to listen and serve webhook server: %v\n", err)
				os.Exit(-1)
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

func (s *webServerImpl) PreConfiguration() {
	s.StringVar(&s.mode, "mode", "", "mode value is https / http")
	s.UintVar(&s.Params.Port, "port", 0, "port using listing service")
	s.X509KeyPairVar(&s.Params.Certs, "tls", "certificado server")
}

func (s *webServerImpl) PostConfiguration() error {
	if commons.StringIsRequiered(&s.mode) != nil {
		return errors.New("mode is requeried")
	}

	vmode, err := commons.ParseMode(&s.mode)
	s.Params.Mode = vmode
	if err != nil {
		return errors.New("mode value must be http/https")
	}

	if *vmode == commons.TLS {
		if s.Params.Certs.PrivateKey == nil {
			return errors.New("no se puedo cargar el certificado 2")
		}
	}
	return nil
}
