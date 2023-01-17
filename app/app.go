package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"project/app/SqlServer"
	"project/app/server"
	"project/pkg/Database"
	"project/pkg/handler"
	"project/pkg/logger"
	"project/pkg/store/sqlBd"
	"syscall"
	"time"
)

// ```
const (
	dbName   = "ww3y_34"
	SOURCE   = "ServerData\\wds_44yy50-252982525.sql"
	PORT     = "3309"
	HOST     = "127.0.0.1"
	USER     = "root"
	PASSWORD = "root"
	Version  = "8.14 gorgon medusa"
)

func Run(config *server.Config, db Database.Database) {
	cfgMLR := server.NewMysqliConfig(dbName, HOST, 3309, Version, nil, nil)
	go func() {
		err := SqlServer.Start(cfgMLR)
		if err != nil {
			logger.Error(err)
		}

	}()
	cfgMLRs := server.NewMysqliConfig("kazuality", HOST, 3310, Version, nil, nil)
	go func() {
		err := SqlServer.Start(cfgMLRs)
		if err != nil {
			logger.Error(err)
		}

	}()
	// Виконуємо команду завантаження дампу у MySQL
	//cmd := exec.Command("mysql", "-u", "root", "-proot", "ww3y-34", "<", "app/SqlServer/Serverdata/wds_44yy50-252982525.sql")
	cmd := exec.Command("mysql", "--host="+HOST, "--port="+PORT, "--password="+PASSWORD, "--user="+USER, dbName, "-e", "source "+SOURCE)
	//fmt.Println(cmd)
	// Записуємо результат виконання команди у буфер
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//fmt.Print(out.Cap())
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	t := time.Now()

	srv := new(server.Server)

	database, err := db.Open(config)
	if err != nil {
		logger.Error(err)

		return
	}
	defer database.Close()

	// init bd
	var store *sqlBd.Store
	store = sqlBd.New(database)

	// init handler
	had := handler.NewHandler(store)

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
