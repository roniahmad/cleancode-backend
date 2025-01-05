package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"wetees.com/api/route"
	"wetees.com/bootstrap"
)

func main() {
	runtime.GOMAXPROCS(2)

	app := fiber.New()
	boot := bootstrap.App()

	var (
		conf   = boot.Conf
		db     = boot.Db
		logger = boot.Log
	)

	defer boot.CloseDBConnection()

	//TODO: CQRS, Rate Limiter

	timeout := time.Duration(conf.ContextTimeout) * time.Second

	route.Setup(db, app, conf, logger, timeout)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", conf.ServicePort)); err != nil {
			logger.Panic().Msg(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // This blocks the main thread until an interrupt is received
	logger.Info().Msg("Gracefully shutting down...")

	_ = app.Shutdown()
	logger.Info().Msgf("%s was successful shutdown", conf.AppName)

}
