package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	//structured login
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	h := api.mount()
	err := api.run(h)
	if err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
