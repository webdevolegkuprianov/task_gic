package task_domain

import (
	"context"
	"time"
)

type IDomain interface {
	GetTaskInfo(ctx context.Context, data []byte) string
	SaveDuration(ctx context.Context, startQueryTime time.Time, endQueryTime time.Time) (err error)
}
