package task_domain

import "context"

type iDao interface {
	getTaskInfo(ctx context.Context) string
}
