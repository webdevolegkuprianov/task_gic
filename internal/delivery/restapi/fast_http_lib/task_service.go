package restapi_fasthttp

import (
	"context"
	"encoding/json"
	"task/internal/view"
	"time"

	"github.com/valyala/fasthttp"
)

type taskServ struct {
	ctx context.Context
	v   *view.View
}

func (serv *taskServ) handleGetTask(next func(ctxHttp *fasthttp.RequestCtx)) func(ctxHttp *fasthttp.RequestCtx) {

	fn := func(ctxHttp *fasthttp.RequestCtx) {

		defer next(ctxHttp)

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		result := serv.v.Views.Task.GetTaskInfo(ctx, ctxHttp.Request.Body())

		dataBin, _ := json.Marshal(result)

		ctxHttp.SetStatusCode(200)
		ctxHttp.SetBody(dataBin)

	}

	return fn

}

func (serv *taskServ) duration() func(ctxHttp *fasthttp.RequestCtx) {

	return func(ctxHttp *fasthttp.RequestCtx) {

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		startTime := ctxHttp.Time()

		endTime := time.Now()

		if err := serv.v.Views.Task.SaveDuration(ctx, startTime, endTime); err != nil {
			ctxHttp.SetStatusCode(500)
			ctxHttp.SetBody([]uint8(err.Error()))
		}

	}

}
