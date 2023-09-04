package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	restapi "task/internal/delivery"
	"task/internal/domains/task_domain"
	"task/internal/view"
)

var (
	l      *log.Logger  = log.Default()
	once   *sync.Once   = new(sync.Once)
	server *http.Server = &http.Server{}
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	once.Do(func() {

		restapi.RegisterRestService(
			ctx,
			view.NewView(
				task_domain.NewDomain(),
			),
			server,
		)

	})

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	defer func() {

		l.Println("отключаю все сервисы приложения")

		signal.Stop(sigCh)

		close(sigCh)

	}()

	go func() {
		l.Println("запускаю http server")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			l.Fatal("не могу запустить HTTP сервер", err)
		}
	}()

	for {
		select {
		case <-sigCh:
			return
		case <-ctx.Done():
			return
		}
	}

}
