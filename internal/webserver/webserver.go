package webserver

import (
	"context"
	"github/ShvetsovYura/pkafka_connect/internal/logger"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Webserver struct {
	server *http.Server
}

func NewWebserver(address string, router *chi.Mux) *Webserver {
	return &Webserver{
		server: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}
}
func (s *Webserver) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	go s.server.ListenAndServe()

	<-ctx.Done()
	logger.Log.Info("Останавливается вэбсервер")
	s.server.Shutdown(ctx)
}
