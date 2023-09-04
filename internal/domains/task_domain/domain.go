package task_domain

import "context"

type Domain struct {
	dao struct {
		task iDao
	}
}

func NewDomain() *Domain {

	return &Domain{
		dao: struct {
			task iDao
		}{
			task: newDao(),
		},
	}

}

func (d *Domain) GetTaskInfo(ctx context.Context) string {

	context, cancel := context.WithCancel(ctx)
	defer cancel()

	return d.dao.task.getTaskInfo(context)
}
