package restapi

import (
	"context"
	"encoding/json"
	"net/http"
	"task/internal/view"
)

type taskServ struct {
	ctx context.Context
	v   *view.View
}

func (serv *taskServ) handleGetTask() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithCancel(serv.ctx)
		defer cancel()

		result := serv.v.Views.Task.GetTaskInfo(ctx)

		dataBin, _ := json.Marshal(result)

		w.WriteHeader(http.StatusOK)
		w.Write(dataBin)

	}

}
