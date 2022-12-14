package bootstrap

import (
	"context"
	"crypto/tls"
	"gogenerate/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/crypto/acme/autocert"
)

const idleTimeout = 5 * time.Second

// Dispatch is handle routing
func Dispatch(ctx context.Context) {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(
		cors.New(),
		logger.New(),
		recover.New(),
	)

	app.Static("/storage/", "/app/public/")
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(struct{ Status string }{"PromptPay is working"})
	})

	// app.Use(pprof.New())
	routes.RegisterRoute(app, ctx)

	/* go func() {
		if errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT")); errApp != nil {
			log.Println("%s", errApp)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// sqlHandler.Close()
	// redisHandler.Close()
	log.Println("Fiber was successful shutdown.")

	// if errApp != nil {
	// 	log.LogError("%s", errApp)
	// } */
	if os.Getenv("DEPLOYMENT_ENV") == "production" {
		// reference: https://github.com/gofiber/recipes/blob/master/https-tls/main.go

		// Certificate manager
		autocertManager := &autocert.Manager{
			// Folder to store the certificates
			Cache: autocert.DirCache("/var/www/.cache"),
		}

		// TLS Config
		tlsConfig := &tls.Config{
			// Get Certificate from Let's Encrypt
			GetCertificate: autocertManager.GetCertificate,
			// By default NextProtos contains the "h2"
			// This has to be removed since Fasthttp does not support HTTP/2
			// Or it will cause a flood of PRI method logs
			// http://webconcepts.info/concepts/http-method/PRI
			NextProtos: []string{
				"http/1.1", "acme-tls/1",
			},
		}

		// Create custom listener
		ln, err := tls.Listen("tcp", os.Getenv("APP_PORT"), tlsConfig)
		if err != nil {
			panic(err)
		}

		log.Fatal(app.Listener(ln))
	} else {
		log.Fatal(app.Listen(os.Getenv("APP_PORT")))
	}
}
