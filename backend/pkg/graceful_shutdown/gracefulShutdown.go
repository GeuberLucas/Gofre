package gracefulshutdown

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
)

// var waitGroup sync.WaitGroup
type GracefulShutdown struct {
	db         *sql.DB
	broker     messaging.IMessaging
	httpServer *http.Server
}

func NewGracefulShutdown(db *sql.DB,
	messagingBroker messaging.IMessaging,
	httpServer *http.Server) *GracefulShutdown {
	return &GracefulShutdown{
		db: db, broker: messagingBroker, httpServer: httpServer,
	}
}
func (gs *GracefulShutdown) ListenSignals() {
	quit:= make(chan os.Signal,1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := gs.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Http server Shutdown with error: %v \n", err)
	} else {
		log.Printf("Http server Shutdown success \n")
	}

	if gs.db != nil {
		if err := gs.db.Close(); err != nil {
			log.Printf("Database connection Shutdown with error: %v", err)
		} else {
			log.Println("Database connection Shutdown success.")
		}
	}
	if gs.broker != nil {
		if err := gs.broker.CloseConnection(); err != nil {
			log.Printf("Broker connection Shutdown with error: %v", err)
		} else {
			log.Println("Broker connection Shutdown success.")
		}
	}

	log.Printf("Service now is down, thanks for using this application")
}
