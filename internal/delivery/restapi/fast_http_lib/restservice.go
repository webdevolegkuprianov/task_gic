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

}

func (s *restService) initRouting() {

	v1 := s.router.Group("/api/v1")

	{
		tasks := v1.Group("/tasks_fast_http")
		tasks.GET("/", s.collection.ts.handleGetTask())
	}

}
