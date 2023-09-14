package restapi_net_http

import (
	"context"
	"fmt"
	"net/http"
	"task/internal/view"

	"github.com/gorilla/mux"
)

type restService struct {
	collection struct {
		ts *taskServ
	}
	router *mux.Router
}

func RegisterRestService(ctx context.Context, v *view.View, srv *http.Server, port string) {

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
		router: mux.NewRouter(),
	}

	rs.initRouting()

	srv.Addr = port
	srv.Handler = rs

}

func (s *restService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *restService) initRouting() {

	v1 := s.router.PathPrefix("/api/v1").Subrouter()

	v1.Use(s.collection.ts.middelwareCounter)

	{
		tasks := v1.PathPrefix("/tasks_net_http").Subrouter()
		tasks.HandleFunc("", s.collection.ts.handleGetTask()).Methods(http.MethodGet)
	}

}
