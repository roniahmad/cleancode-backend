package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"wetees.com/api/route"
	"wetees.com/bootstrap"
	"wetees.com/internal/vars"
)

func main() {
	runtime.GOMAXPROCS(2)

	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	boot := bootstrap.App()

	var (
		conf      = boot.Conf
		db        = boot.Db
		logger    = boot.Log
		rateLimit = conf.RateLimit
	)

	defer boot.CloseDBConnection()

	// limit N requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        rateLimit,
	}))

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: conf.Mode == vars.ModeProd,
	}))

	timeout := time.Duration(conf.ContextTimeout) * time.Second

	route.Setup(db, app, conf, logger, timeout)

	go func() {
		if err := app.Listen(fmt.Sprintf("0.0.0.0:%s", conf.ServicePort)); err != nil {
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
