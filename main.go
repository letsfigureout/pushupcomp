package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/letsfigureout/pushupcomp/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := log.New(os.Stdout, "goapp: ", log.LstdFlags|log.Ltime|log.Lshortfile)

	// Local environment variables
	godotenv.Load()

	if err := run(log); err != nil {
		log.Fatalf("fatal: %v", err)
	}

}

func run(log *log.Logger) error {
	// Shutdown Signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	conf := struct {
		ShutdownTimeout time.Duration
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
	}{
		ShutdownTimeout: 30 * time.Second,
		ReadTimeout:     30 * time.Second,
		WriteTimeout:    30 * time.Second,
	}

	/*
	Setup Database
	 */
	dbConn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer dbConn.Close(context.Background())


	/*
	  Setup API
	 */
	// Listen port from config_vars
	appPort := os.Getenv("PORT")
	if appPort == "" {
		return fmt.Errorf("PORT environment variable not set")
	}

	server := http.Server{
		Addr:    ":" + appPort,
		Handler: routes.API(log),
	}
	serverErrors := make(chan error, 1)

	go func() {

		log.Printf("API Listening on :%v\n", appPort)
		serverErrors <- server.ListenAndServe()

	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error : %v", err)

	case sig := <-shutdown:
		log.Println("shutdown signal deteted, attempting graceful shutdown...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not shutdown gracefully : %v", err)
		}

	}

	return nil
}
