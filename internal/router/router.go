package router

import (
	"github/ShvetsovYura/pkafka_connect/internal/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Renderer interface {
	Render(metrics []types.Metric) ([]byte, error)
}
type MetricListGetter interface {
	GetList() []types.Metric
}

type Router struct {
	tmpl    Renderer
	metrics MetricListGetter
	router  *chi.Mux
}

func NewRouter(tplSvc Renderer, metricsSvc MetricListGetter) *Router {
	router := chi.NewRouter()

	r := &Router{tmpl: tplSvc, metrics: metricsSvc, router: router}
	r.router.Get("/metrics", r.MetricHandler)
	return r
}

func (r *Router) GetRouter() *chi.Mux {
	return r.router
}

func (r *Router) MetricHandler(w http.ResponseWriter, req *http.Request) {
	rendered, err := r.tmpl.Render(r.metrics.GetList())
	if err != nil {
		w.WriteHeader(422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(rendered)
}
