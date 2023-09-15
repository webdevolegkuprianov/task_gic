package restapi_net_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task/internal/view"
	"time"
)

type taskServ struct {
	ctx context.Context
	v   *view.View
}

func (serv *taskServ) handleTest1() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		var db []uint8

		db, _ = io.ReadAll(r.Body)

		result := serv.v.Views.Task.GetTaskInfo(ctx, db)

		dataBin, _ := json.Marshal(result)

		w.WriteHeader(http.StatusOK)
		w.Write(dataBin)

	}

}

func (serv *taskServ) handleTest(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		ctxHttp := context.WithValue(r.Context(), "time_start", time.Now())

		r = r.WithContext(ctxHttp)

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		var db []uint8

		db, _ = io.ReadAll(r.Body)

		result := serv.v.Views.Task.GetTaskInfo(ctx, db)

		dataBin, _ := json.Marshal(result)

		w.WriteHeader(http.StatusOK)
		w.Write(dataBin)

		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)

}

func (serv *taskServ) duration() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		endTime := time.Now()

		startTime := r.Context().Value("time_start").(time.Time)

		fmt.Println("startTime: ", startTime)

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		if err := serv.v.Views.Task.SaveDuration(ctx, startTime, endTime); err != nil {
			fmt.Println(err)
		}

	}

}
