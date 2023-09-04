package view

import "task/internal/domains/task_domain"

type View struct {
	Views struct {
		Task *TaskView
	}
}

func NewView(domTask *task_domain.Domain) *View {

	return &View{
		Views: struct {
			Task *TaskView
		}{
			Task: &TaskView{
				dom: struct {
					task task_domain.IDomain
				}{
					task: domTask,
				},
			},
		},
	}

}
