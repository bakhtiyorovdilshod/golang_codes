package main 

import (
	"fmt"
)
type Task struct {
	assignee string
	title string
	time string
	done bool

}

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() *TaskManager {
 	tm := TaskManager{}
 	tm.tasks = []Task{}

 	return &tm
}

func (task *TaskManager) add_task(item Task) {
    task.tasks = append(task.tasks, item)
}

func (task *TaskManager) list_task() []Task{
	return task.tasks
}

func (task * TaskManager) update_task(id int, t Task) {
	task.tasks[id] = t
}

func (task * TaskManager) delete_task(id int) {
	var i int 
	i = id
	task.tasks = append(task.tasks[:i], task.tasks[i+1:]...)

}

func main() {
	tm := NewTaskManager()
	tm.add_task(Task{assignee:"dilshod", title:"du", time:"2017-09-12",done:true})
	x:=tm.list_task()[0]
	fmt.Println(x)
   

}
