package main

import "log/slog"

func main() {
	app, err := app.New()
	if err != nil {
		slog.Error("failed to create app", "error", err)
		return
	}
	if err = app.Run(); err != nil {
		slog.Error("failed to run app", "error", err)
		return
	}
}
