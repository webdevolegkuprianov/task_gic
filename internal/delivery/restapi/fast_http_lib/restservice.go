package restapi_fasthttp

import (
	"context"
	"fmt"
	"task/internal/view"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type restService struct {
	collection struct {
		ts *taskServ
	}
	router *router.Router
}

func RegisterRestService(ctx context.Context, v *view.View, srv *fasthttp.Server, port string) {

	if v == nil {
		panic(fmt.Errorf("не инициализирован View"))
	}

	rs := &restService{
		collection: struct {
			ts *taskServ
		}{
			ts: &taskServ{
				ctx: ctx,
				v:   v,
			},
		},
		router: router.New(),
	}

	rs.initRouting()

	srv.Handler = rs.router.Handler
	srv.Concurrency = 5000
	srv.ReadBufferSize = 2 * 1024
	srv.MaxRequestBodySize = 1024 * 1024

}

func (s *restService) initRouting() {

	v1 := s.router.Group("/api/v1")

	{
		tasks := v1.Group("/fast_http")
		tasks.POST("/", s.collection.ts.handleGetTask(s.collection.ts.duration()))
	}

}

/*
func (s *restService) recovery(next func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {

	fn := func(ctx *fasthttp.RequestCtx) {

		defer func() {
			if rvr := recover(); rvr != nil {
				var msg string
				switch v := rvr.(type) {
				case error:
					msg = v.Error()
				case string:
					msg = v
				default:
					msg = fmt.Sprint(v)
				}
				ctx.Error(msg, 500)
			}
		}()

		next(ctx)
	}

	return fn

}
*/
