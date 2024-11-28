package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO       TaskStatus = "todo"
	TASK_STATUS_INPROGRESS TaskStatus = "in-process"
	TASK_STATUS_DONE       TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(list_limited string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}
	for _, task := range tasks {
		if list_limited == "" || string(task.Status) == list_limited {
			fmt.Printf("%d. %s\n", task.ID, task.Description)
		}
	}
	return nil
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	id := int64(len(tasks) + 1)
	newTask := NewTask(id, description)
	tasks = append(tasks, *newTask)
	err = WriteTasksToFile(tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}

func FindTask(id int64) (*Task, error) {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("Task with ID %d not found", id)
}

func UpdateTask(id int64, description string) error {
	task, err := FindTask(id)
	if err != nil {
		return err
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = *task
		}
	}
	err = WriteTasksToFile(tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Task updated successfully (ID :%d)\n", id)
	return nil
}

func MarkTask(id int64, status TaskStatus) error {
	task, err := FindTask(id)
	if err != nil {
		return err
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = *task
		}
	}
	err = WriteTasksToFile(tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Task marked as %s (ID: %d)\n", status, id)
	return nil
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err = WriteTasksToFile(tasks)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("Task deleted successfully (ID: %d)\n", id)
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}
