package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		time.Sleep(10 * time.Second)
		return c.String(http.StatusOK, "Server is running \n")
	})

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatal("Shutting down our server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1) // 1 isliye diya kyuki koi bhi signal miss na ho

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit // program will wait here and will get block until we get shutdown signal, jab tak shutdown ka notification nahi ata

	log.Println("\n Shutdown signal is received")

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second) //forever waiting stage me naa jaye isliye
	defer cancel()

	//graceful shutdown:- nayi request accept nahi krna hai

	if err := e.Shutdown(ctx)
	err != nil {
		log.Fatal("Server forced shutdown", err)
	}

	log.Println("Server exited gracefully")
}
//til: Idiomatic Go prefers if err := ...; 
//err != nil when the error is only needed locally, and a separate err := when the error must be reused.