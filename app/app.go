package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"project/app/server"
	"project/pkg/handler"
	"project/pkg/logger"
	"project/pkg/store/sqlBd"
	"syscall"
	"time"
)

func Run(config *server.Config) {
	srv := new(server.Server)
	db, err := newDB(config)
	if err != nil {
		logger.Error(err)

		return
	}
	defer db.Close()
	t := time.Now()

	// init bd
	storet := make(chan *sqlBd.Store, 1)
	go func() {
		storet <- sqlBd.New(db)
	}()

	// init handler
	had := handler.NewHandler(<-storet)

	go func() {
		if err := srv.Run(config, had); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Infof("Server started, second %.2f", time.Since(t).Seconds())

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Block the end
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

}
func newDB(c *server.Config) (*sql.DB, error) {
	db, err := sql.Open(c.DriverName, c.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	return db, nil
}
