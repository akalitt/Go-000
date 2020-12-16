package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github/akalitt/go-errors-example/config"
	"github/akalitt/go-errors-example/internal/model"
	"github/akalitt/go-errors-example/internal/server/router"
	"github/akalitt/go-errors-example/middleware"
	"github/akalitt/go-errors-example/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	config.Parse()

	model.Database()

	logger.Init("INFO")

	g := gin.Default()

	router.Load(
		// Cores.
		g,
		middleware.Logging(),
	)

	model.ResetDatabase()

	srv := &http.Server{
		Addr:    viper.GetString("port"),
		Handler: g,
	}

	go func() {
		// service connections
		fmt.Println("server Start AT:", viper.GetString("port"))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefulShutdown(srv)
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server Exiting")
}
