package main

import "testing"

var (
	tm *TaskManager

)


func TestNewTaskManagerAdd (t *testing.T) {
	tm = NewTaskManager()
	tm.add_task(Task{assignee:"dilshod", title:"du", time:"2017-09-12",done:true})
	if len(tm.tasks)== 0 {
		t.Error("Task not added")
	}
}

func TestTaskManagerUpdate (t *testing.T) {
	tm = NewTaskManager()
	tm.add_task(Task{assignee:"dilshod", title:"du", time:"2017-09-12",done:true})
	x:=tm.list_task()[0]
	tm.update_task(0,Task{assignee:"dilshodbek", title:"du", time:"2017-09-12",done:true})
	if tm.list_task()[0]==x {
		t.Error("Error")

	}
}

func TestTaskManagerDelete (t *testing.T) {
	tm = NewTaskManager()
	tm.add_task(Task{assignee:"dilshod", title:"du", time:"2017-09-12",done:true})
	my_task_len:= len(tm.tasks)
	tm.delete_task(0)
	check_task_len:=len(tm.tasks)
	if my_task_len == check_task_len {
		t.Error("Error")

	}

}
