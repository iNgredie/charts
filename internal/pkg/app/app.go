package app

import (
	"errors"
	"log/slog"
	"net/http"
)

type App struct {
	logger *slog.Logger

	dbConn *pgxpool.Pool
}

func New() (*App, error) {
	var err error
	app := &App{}
	return app, err
}

func (a *App) Run() error {
	go func() {
		err := a.http.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("http server failed", "error", err)
		}
	}()
	<-a.closeCh
	for i := len(a.closers) - 1; i >= 0; i-- {
		err := a.closers[i]()
		if err != nil {
			a.logger.Error("failed to close resource", "i", i, "error", err)
		}
	}
	return nil
}

func (a *App) initLogger() {
	logger := slog.Default()
	a.logger = logger
}
