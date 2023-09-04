package task_domain

import "context"

type dao struct {
	taskInfo string
}

func newDao() *dao {
	return &dao{
		taskInfo: "task",
	}
}

func (dao *dao) getTaskInfo(ctx context.Context) string {
	return dao.taskInfo
}
