package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.Int("addr", 8000, "HTTP network address")
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	addrStr := fmt.Sprintf(":%d", *addr)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error("failed to connect to database", slog.String("dsn", *dsn), slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("starting server", slog.String("port", addrStr))

	if err = http.ListenAndServe(addrStr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
