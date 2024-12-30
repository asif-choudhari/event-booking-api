package main

import (
	"context"
	"event-booking-api/db"
	"event-booking-api/middlewares"
	"event-booking-api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set release mode before creating router
	gin.SetMode(gin.ReleaseMode)

	db.InitDb()
	router := gin.Default()

	router.GET("/events", routes.GetAllEvents)
	router.GET("/event/:id", routes.GetEvent)
	router.POST("/user/signup", routes.SignUpUser)
	router.POST("/user/login", routes.LoginUser)

	protected := router.Group("/")
	protected.Use(middlewares.Authenticate)
	protected.POST("/event", middlewares.Authenticate, routes.AddEvent)
	protected.PUT("/event", routes.UpdateEvent)
	protected.DELETE("/event/:id", routes.DeleteEvent)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s\n", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
