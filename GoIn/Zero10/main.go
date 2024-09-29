package main

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

func (t Task) IsTopPriority() bool {
    return t.priority == 4 || t.priority == 5
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

func (t *ToDoList) TasksCount() int {
    return len(t.tasks)
}

func (t *ToDoList) NotesCount() int {
    return len(t.notes)
}

func (t *ToDoList) CountTopPrioritiesTasks() int {
    count := 0
    for _, task := range t.tasks {
        if task.IsTopPriority() {
            count++
        }
    }
    return count
}

func (t *ToDoList) CountOverdueTasks() int {
    count := 0
    for _, task := range t.tasks {
        if task.IsOverdue() {
            count++
        }
    }
    return count
}


func main() {
	todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}, Task{summary: "Make Yandex Lyceum Task 10", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 3}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
	fmt.Println(todo.TasksCount())
	fmt.Println(todo.NotesCount())
	fmt.Println(todo.CountTopPrioritiesTasks())
	fmt.Print(todo.CountOverdueTasks())
}