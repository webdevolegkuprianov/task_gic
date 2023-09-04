package view

import (
	"context"
	"task/internal/domains/task_domain"
)

type TaskView struct {
	dom struct {
		task task_domain.IDomain
	}
}

func (view *TaskView) GetTaskInfo(ctx context.Context) (result *TaskInfo) {

	result = &TaskInfo{
		Data: view.dom.task.GetTaskInfo(ctx),
	}

	return

}
