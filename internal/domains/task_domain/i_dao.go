package task_domain

import "context"

type iDao interface {
	getTaskInfo(ctx context.Context, data []byte) string
	updateDurationFile(ctx context.Context, duration int64) (err error)
}
