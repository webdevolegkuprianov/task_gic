package task_domain

import "context"

type iDao interface {
	getTaskInfo(ctx context.Context) string
	updateTaskFile(ctx context.Context, value int) (err error)
}
