package http

import (
	"errors"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Srv *http.Server
}

func NewServer(cfgServer config.Server, handler http.Handler) *Server {
	addr := fmt.Sprintf("%s:%d", cfgServer.Host, cfgServer.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{
		srv,
	}
}

func (s *Server) Start() error {
	log.Printf("http server listening at %s", s.Srv.Addr)
	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("error starting http server:", err)
		}
	}()

	return nil
}

func (s *Server) Stop() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	log.Printf("server successfully stopped received signal %s", sign.String())
}
