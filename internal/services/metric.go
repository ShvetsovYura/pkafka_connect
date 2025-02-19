package services

import (
	"context"
	"github/ShvetsovYura/pkafka_connect/internal/logger"
	"github/ShvetsovYura/pkafka_connect/internal/types"
	"log/slog"
)

type MetricsService struct {
	metrics map[string]types.Metric
}

func NewMeticsService() *MetricsService {
	return &MetricsService{
		metrics: make(map[string]types.Metric, 0),
	}
}

func (s *MetricsService) Put(name string, value types.Metric) {
	s.metrics[name] = value
}

func (s *MetricsService) GetList() []types.Metric {
	mList := make([]types.Metric, 0)
	for _, v := range s.metrics {
		mList = append(mList, v)
	}
	return mList
}

func (s *MetricsService) PutMany(metrics map[string]types.Metric) {
	for k, v := range metrics {
		metrics[k] = v
	}
}

func (s *MetricsService) Run(ctx context.Context, metrics <-chan types.Metric) {
	for {
		select {
		case <-ctx.Done():
			logger.Log.Info("Останавливается обработка метрик")
			return
		case m := <-metrics:
			logger.Log.Info("proccess messages", slog.Int("queue", len(metrics)))
			s.Put(m.Name, m)
		}
	}
}
