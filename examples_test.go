package chslog_test

import (
	slogch "atomicgo.dev/chslog"
	"log/slog"
	"os"
)

func ExampleChoose() {
	// Uses a text logger in development and a JSON logger in production.
	prodHandler := slog.NewJSONHandler(os.Stdout, nil)
	devHandler := slog.NewTextHandler(os.Stdout, nil)

	handler := slogch.Choose(prodHandler, devHandler)
	logger := slog.New(handler)

	logger.Info("Hello, World!", "foo", "bar")
}
