package main

import "fmt"

var (
	priorityStore = map[int]int{1: 1, 2: 2}
	riskStore     = map[int]float32{1: 0.2, 2: 0.6}
)

type Task struct {
	id          int
	description string
}

func moreSafe(id1 int, id2 int) int {
	risk1 := riskStore[id1]
	risk2 := riskStore[id1]
	if risk1 > risk2 {
		return id1
	} else {
		return id2
	}
}
func morePriority(id1 int, id2 int) int {
	priority1 := priorityStore[id1]
	priority2 := priorityStore[id1]
	if priority1 > priority2 {
		return id1
	} else {
		return id2
	}
}
func chooseTask(task1 Task, task2 Task, filter func(int, int) int) Task {
	id := filter(task1.id, task2.id)
	if id == task1.id {
		return task1
	} else {
		return task2
	}

}
func main() {
	task1 := Task{id: 1, description: "Create a New Logo"}
	task2 := Task{id: 2, description: "Migrate a DB"}
	morePriorityTask := chooseTask(task1, task2, morePriority)
	moreRiskyTask := chooseTask(task1, task2, moreSafe)
	fmt.Println("More Priority tasks:", morePriorityTask)
	fmt.Println("Tasks with big bussiness risks:", moreRiskyTask)

}
