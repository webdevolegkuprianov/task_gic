package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"task/internal/domains/task_domain"
	"task/internal/view"
	"task/pkg/configs"

	restapi_fasthttp "task/internal/delivery/restapi/fast_http_lib"
	restapi_net_http "task/internal/delivery/restapi/net_http_lib"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

var (
	l              *log.Logger      = log.Default()
	once           *sync.Once       = new(sync.Once)
	netHttpServer  *http.Server     = &http.Server{}
	fastHttpServer *fasthttp.Server = &fasthttp.Server{}
	envCnf         *configs.Config
	chCounter      chan (int) = make(chan int)
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

		restapi_net_http.RegisterRestService(
			ctx,
			view.NewView(
				task_domain.NewDomain(envCnf.FilePath, chCounter),
			),
			netHttpServer,
			envCnf.Port1,
		)

		restapi_fasthttp.RegisterRestService(
			ctx,
			view.NewView(
				task_domain.NewDomain(envCnf.FilePath, chCounter),
			),
			fastHttpServer,
			envCnf.Port2,
		)

		/*
			stdservice.RegisterStdService(
				ctx,
				chCounter,
			)
		*/

	})

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	defer func() {

		l.Println("отключаю все сервисы приложения")
		signal.Stop(sigCh)
		close(sigCh)
		close(chCounter)

	}()

	var ln net.Listener

	ln, err = reuseport.Listen("tcp4", envCnf.Port2)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := netHttpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			l.Fatal("не могу запустить HTTP сервер", err)
		}
	}()

	go func() {
		if err := fastHttpServer.Serve(ln); !errors.Is(err, http.ErrServerClosed) {
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
