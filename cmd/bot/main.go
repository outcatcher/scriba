package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/outcatcher/scriba/internal/bot"
	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/usecases"
)

var configPath = flag.String("config", "./config/config.yaml", "Configuration file to be used")

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath, true)
	if err != nil {
		log.Fatalln("failed to load configuration", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	app, err := usecases.InitApp(cfg)
	if err != nil {
		log.Fatalln("error initializing app", err)
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		// waiting for signal
		sig := <-sigChan
		slog.Error("received signal:", "signal", sig)
		cancel()
	}()

	if err := bot.Start(ctx, cfg.Bot, app); err != nil {
		log.Fatalln("failed to start bot", err)
	}
}
