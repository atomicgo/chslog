/*
Package chslog provides a simple way to set up a logger that automatically chooses between a production and a development logger.

Example:

	func main() {
		// Uses a text logger in development and a JSON logger in production.
		prodHandler := slog.NewJSONHandler(os.Stdout, nil)
		devHandler := slog.NewTextHandler(os.Stdout, nil)

		handler := slogch.Choose(prodHandler, devHandler)
		logger := slog.New(handler)

		logger.Info("Hello, World!", "foo", "bar")
		// Prod: {"time":"2023-08-03T01:31:27.6681464+02:00","level":"INFO","msg":"Hello, World!","foo":"bar"}
		// Dev:  time=2023-08-03T01:30:23.438+02:00 level=INFO msg="Hello, World!" foo=bar
	}
*/
package chslog
