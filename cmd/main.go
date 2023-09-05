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

	"task/internal/delivery/restapi"
	stdservice "task/internal/delivery/stdout"
	"task/internal/domains/task_domain"
	"task/internal/view"
	"task/pkg/configs"
)

var (
	l         *log.Logger  = log.Default()
	once      *sync.Once   = new(sync.Once)
	server    *http.Server = &http.Server{}
	envCnf    *configs.Config
	chCounter chan (int) = make(chan int)
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

	envCnf, err = configs.NewConfig()
	if err != nil {
		l.Fatalf("envconfig fail, %s", err.Error())
	}

	once.Do(func() {

		restapi.RegisterRestService(
			ctx,
			view.NewView(
				task_domain.NewDomain(envCnf.FilePath, chCounter),
			),
			server,
			envCnf.Port,
		)

		stdservice.RegisterRestService(
			ctx,
			chCounter,
		)

	})

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	defer func() {

		l.Println("отключаю все сервисы приложения")
		signal.Stop(sigCh)
		close(sigCh)
		close(chCounter)

	}()

	go func() {
		l.Printf("запускаю http server на порту %s\n", envCnf.Port)
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
