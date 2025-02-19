package main

import (
	"context"
	"github/ShvetsovYura/pkafka_connect/internal/consumer"
	"github/ShvetsovYura/pkafka_connect/internal/logger"
	"github/ShvetsovYura/pkafka_connect/internal/router"
	"github/ShvetsovYura/pkafka_connect/internal/services"
	"github/ShvetsovYura/pkafka_connect/internal/types"
	"github/ShvetsovYura/pkafka_connect/internal/webserver"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gopkg.in/yaml.v3"
)

func main() {
	logger.Init()
	data, err := os.ReadFile("config.yml")
	if err != nil {
		logger.Log.Warn(err.Error())
		return
	}

	var config types.Options

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}

	logger.Log.Info("Запуск с конфигом: ", slog.Any("cfg", config))
	queue := make(chan types.Metric, 1000)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		close(queue)
		stop()
	}()

	t, err := services.NewTemplateService()
	if err != nil {
		log.Fatal(err)
	}

	c := consumer.NewKafkaConsumer(config.Consumer)
	m := services.NewMeticsService()
	r := router.NewRouter(t, m)

	s := webserver.NewWebserver(config.Webserver.Address, r.GetRouter())
	go c.Run(ctx, queue)
	go m.Run(ctx, queue)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go s.Run(ctx, wg)
	wg.Wait()
}
