package ex5

import (
	"fmt"
	"time"
)

type Task struct {
    summary     string
    description string
    deadline    time.Time
    priority    int
}

func (t Task) IsOverdue() bool {
    return time.Now().After(t.deadline)
}


func (t Task) Print() {
	fmt.Printf("Title: %s\n", t.summary)
	fmt.Printf("Description: %s\n", t.description)
	fmt.Printf("Deadline: %s\n", t.deadline.Format("2006-01-02"))
	fmt.Printf("Priority: %t\n", t.priority)
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (todolist ToDoList) TasksCount() int {
	return len(todolist.tasks)
}

func (todolist ToDoList) NotesCount() int {
	return len(todolist.notes)
}

func (t Task) IsTopPriority() bool {
    return t.priority == 4 || t.priority == 5
}


func (todolist ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range todolist.tasks {
		if task.IsTopPriority() {
			count++
		}
	}
	return count
}

func (todolist ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range todolist.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}

