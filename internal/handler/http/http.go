package http

import (
	"errors"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Srv *http.Server
	Log *zap.SugaredLogger
}

func NewServer(cfgServer config.Server, handler http.Handler, log *zap.SugaredLogger) *Server {
	addr := fmt.Sprintf("%s:%d", cfgServer.Host, cfgServer.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{
		Srv: srv,
		Log: log,
	}
}

func (s *Server) Start() error {
	s.Log.Infof("http server listening at %s", s.Srv.Addr)

	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Log.Fatal("error starting http server:", err)
		}
	}()

	return nil
}

func (s *Server) Stop() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	s.Log.Infof("server successfully stopped received signal %s", sign.String())
}
