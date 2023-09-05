package task_domain

import "context"

type IDomain interface {
	GetTaskInfo(ctx context.Context) string
	IncrementCounter(ctx context.Context) (err error)
}
