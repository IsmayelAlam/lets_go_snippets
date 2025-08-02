package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.Int("addr", 4000, "HTTP network address")
	flag.Parse()
	addrStr := fmt.Sprintf(":%d", *addr)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("port", addrStr))

	if err := http.ListenAndServe(addrStr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
