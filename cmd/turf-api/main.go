package main

import (
	"context"
	"expvar"
	"log"
	"net/http"
	_ "net/http/pprof" // Register the pprof handlers
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tomchavakis/turf-api/config"
	phhtp "github.com/tomchavakis/turf-api/internal/infra/http"
	measurement "github.com/tomchavakis/turf-api/internal/infra/repository/geo"
)

var build = "develop"

func main() {
	log := log.New(os.Stdout, "turf-api ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("main: error:", err)
		log.Println("main: Completed")
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	ctx := context.Background()
	cfg := config.New(ctx)

	// App Starting
	expvar.NewString("build").Set(build)
	log.Printf("main: Started : Application initializing : version %q", build)

	// Start API Service
	log.Println("main: Initializing API")

	// channel to listen for an interrupt or terminate singal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)

	msrSvc, err := measurement.New()
	if err != nil {
		log.Printf("error: %v", err)
		return errors.New("main: can't initialize measurement service")
	}

	// HTTP initialisation
	r := phhtp.New(msrSvc)
	r.RouteBuilder()

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      r.Router,
		ReadTimeout:  cfg.Web.APIReadTimeout,
		WriteTimeout: cfg.Web.APIWriteTimeout,
	}

	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Start Debug Service
	// /debug/pprof
	// /debug/vars

	if cfg.Web.DebugMode {
		log.Println("main: Initializing debugging support")
		go func() {
			log.Printf("main: Debug Listening %s", cfg.Web.DebugHost)
			if err := http.ListenAndServe(cfg.Web.DebugHost, http.DefaultServeMux); err != nil {
				log.Printf("main: Debug Listener closed : %v", err)
			}
		}()
	}

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")
	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)
		// Give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = api.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
	return nil
}
