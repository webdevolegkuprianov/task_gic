package task_domain

import (
	"context"
	"time"
)

type Domain struct {
	dao struct {
		task iDao
	}
}

func NewDomain(filePath string, ch chan int) *Domain {

	return &Domain{
		dao: struct {
			task iDao
		}{
			task: newDao(filePath, ch),
		},
	}

}

func (d *Domain) GetTaskInfo(ctx context.Context, data []byte) string {

	context, cancel := context.WithCancel(ctx)
	defer cancel()

	return d.dao.task.getTaskInfo(context, data)
}

func (d *Domain) SaveDuration(ctx context.Context, startQueryTime time.Time, endQueryTime time.Time) (err error) {

	duration := endQueryTime.UnixNano() - startQueryTime.UnixNano()

	if err = d.dao.task.updateDurationFile(ctx, duration); err != nil {
		return
	}

	return
}
