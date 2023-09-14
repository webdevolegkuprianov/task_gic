package restapi_fasthttp

import (
	"context"
	"encoding/json"
	"net/http"
	"task/internal/view"

	"github.com/valyala/fasthttp"
)

type taskServ struct {
	ctx context.Context
	v   *view.View
}

func (serv *taskServ) handleGetTask() func(ctx *fasthttp.RequestCtx) {

	return func(ctxHttp *fasthttp.RequestCtx) {

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		result := serv.v.Views.Task.GetTaskInfo(ctx)

		dataBin, _ := json.Marshal(result)

		ctxHttp.SetStatusCode(200)
		ctxHttp.SetBody(dataBin)

	}

}

func (serv *taskServ) middelwareCounter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		serv.v.Views.Task.IncrementCounter(ctx)

		next.ServeHTTP(w, r)

	})

}
