package task

import (
	"encoding/json"
	"fmt"
	"time"
)

type Status string

const (
	TODO        Status = "TODO"
	IN_PROGRESS Status = "IN_PROGRESS"
	DONE        Status = "DONE"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t Task) String() string {
	return fmt.Sprintf(
	"Task Id %d\n"+ 
	"Current Status: %s\n"+
	"Created At: %s\n"+
	"Last Update: %s\n", t.Id, t.Status, t.CreatedAt.Format(time.ANSIC), t.UpdatedAt.Format(time.ANSIC),
	)
}

func AddTask(description string) (*Task, error) {
	data, err := readDataFile()
	if err != nil {
		return nil, err
	}

	var (
		tasks  []*Task
		lastId int
	)

	if len(data) != 0 {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return nil, fmt.Errorf("[!] Error while parsing json data: %w", err)
		}

		lastId = tasks[len(tasks)-1].Id

	}
	newTask := &Task{
		Id:          lastId + 1,
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	encodedTasks, _ := json.MarshalIndent(tasks, "", "	")

	if err = updateDataFile(encodedTasks); err != nil {
		return nil, err
	}

	return newTask, nil

}

func GetTasks(id int) (*Task, []*Task, error) {
	tasks, err := getTasks()
	if err != nil {
		return nil, nil, err
	}

	for _, task := range tasks {
		if task.Id == id {
			return task, tasks, nil
		}
	}
	return nil, nil, fmt.Errorf("[!] Couldn't find a task with provided id")
}

func UpdateTasks(newTask *Task, tasks []*Task) (*Task, error) {
	var modifiedTask *Task

	for _, task := range tasks {
		if task.Id == newTask.Id {
			task = newTask
			newTask.UpdatedAt = time.Now()
			break
		}
	}

	encodedTasks, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return nil, fmt.Errorf("[!] Error while updating data: %w", err)
	}

	if err = updateDataFile(encodedTasks); err != nil {
		return nil, err
	}

	return modifiedTask, nil
}

func DeleteTask(id int) error {
	tasks, err := getTasks()
	if err != nil {
		return err
	}

	for index, task := range tasks {
		if task.Id == id {
			last := len(tasks) - 1
			tasks[index], tasks[last] = tasks[last], tasks[index]
			tasks = tasks[:last]
			break
		}
	}

	encodedTasks, _ := json.MarshalIndent(tasks, "", "	")
	if err = updateDataFile(encodedTasks); err != nil {
		return err
	}

	return nil
}

func ListAll() error {
	tasks, err := getTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(*task)
	}

	return nil
}

func ListSpecificTasks(s Status) error {
	tasks, err := getTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Status == s {
			fmt.Println(task)
		}
	}

	return nil
}

// func ListInProgress() ([]*Task, error) {
// 	tasks, err := getTasks()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var i int
// 	for index, task := range tasks {
// 		if task.Status == IN_PROGRESS {
// 			tasks[i], tasks[index] = tasks[index], tasks[i]
// 		}
// 	}

// 	return tasks[:i], nil
// }

// func ListDone() ([]*Task, error) {
// 	tasks, err := getTasks()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var i int
// 	for index, task := range tasks {
// 		if task.Status == DONE {
// 			tasks[i], tasks[index] = tasks[index], tasks[i]
// 		}
// 	}

// 	return tasks[:i], nil
// }
