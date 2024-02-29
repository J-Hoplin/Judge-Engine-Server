package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"judge-engine/internal/controller"
	"judge-engine/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func ApplicationBootstrap(port int) error {
	var err error
	var server *http.Server

	// Bootstrap of database
	err = database.MySQLBootstrap()
	if err != nil {
		return err
	}
	// Close server with graceful shutdown
	defer database.CloseClient()

	// Set gin colorize log
	gin.ForceConsoleColor()
	r := gin.Default()

	// Enroll Route
	controller.SubmissionRoute(r)

	// Server
	server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}
	// Set graceful shutdown of server
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Println("Server shutdown with signal: " + sig.String())
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Println("Server abnormally shutdown: " + err.Error())
	}
	select {
	case <-ctx.Done():
		log.Println("Server shutdown successfully!")
	}
	return nil
}
