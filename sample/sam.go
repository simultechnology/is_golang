package main

import (
	"fmt"
)

type Task struct {
	ID	int
	Detail	string
	done	bool
}

func (t Task) setDone(done bool) {
	t.done = done
}

func (t *Task) setDoneByPointer(done bool) {
	t.done = done
}

func ChangeDetail(task Task) {
	task.Detail = "fugafuga"
}
func ChangeDetailByPointer(task *Task) {
	task.Detail = "fugafuga"
}

func main()  {
	fmt.Println("start")
	task := Task{
		1, "hogehoge", true,
	}

	fmt.Println(task)
	task.setDone(false)
	fmt.Println(task)
	task.setDoneByPointer(false)
	fmt.Println(task)

	ChangeDetail(task)
	fmt.Println(task)

	fmt.Println("----------------------------")

	task2 := &Task{
		1, "hogehoge", true,
	}
	fmt.Println(task2)
	fmt.Println(task2.Detail)
	ChangeDetailByPointer(task2)
	fmt.Println(task2)
	fmt.Println(task2.Detail)
}

