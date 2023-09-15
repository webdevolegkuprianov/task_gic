package view

import (
	"context"
	"task/internal/domains/task_domain"
	"time"
)

type TaskView struct {
	dom struct {
		task task_domain.IDomain
	}
}

func (view *TaskView) GetTaskInfo(ctx context.Context, data []byte) (result *TaskInfo) {

	result = &TaskInfo{
		Data: view.dom.task.GetTaskInfo(ctx, data),
	}

	return

}

func (view *TaskView) SaveDuration(ctx context.Context, startTime, endTime time.Time) (err error) {

	if err = view.dom.task.SaveDuration(ctx, startTime, endTime); err != nil {
		return
	}

	return

}
