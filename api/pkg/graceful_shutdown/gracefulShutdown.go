package gracefulshutdown

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

// var waitGroup sync.WaitGroup
type GracefulShutdown struct {
	db         *sql.DB
	httpServer *fiber.App
}

func NewGracefulShutdown(db *sql.DB,
	httpServer *fiber.App) *GracefulShutdown {
	return &GracefulShutdown{
		db: db, httpServer: httpServer,
	}
}
func (gs *GracefulShutdown) ListenSignals() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := gs.httpServer.Shutdown(); err != nil {
		log.Printf("Fiber server Shutdown with error: %v \n", err)
	} else {
		log.Printf("Fiber server Shutdown success \n")
	}

	if gs.db != nil {
		if err := gs.db.Close(); err != nil {
			log.Printf("Database connection Shutdown with error: %v", err)
		} else {
			log.Println("Database connection Shutdown success.")
		}
	}

	log.Printf("Service now is down, thanks for using this application")
}
