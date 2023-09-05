package task_domain

import "context"

type Domain struct {
	dao struct {
		task iDao
	}
}

func NewDomain(filePath string) *Domain {

	return &Domain{
		dao: struct {
			task iDao
		}{
			task: newDao(filePath),
		},
	}

}

func (d *Domain) GetTaskInfo(ctx context.Context) string {

	context, cancel := context.WithCancel(ctx)
	defer cancel()

	return d.dao.task.getTaskInfo(context)
}

func (d *Domain) IncrementCounter(ctx context.Context) (err error) {

	if err = d.dao.task.updateTaskFile(ctx, 1); err != nil {
		return
	}

	return
}
