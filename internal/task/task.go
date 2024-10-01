package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	TODO        = "TODO"
	IN_PROGRESS = "IN_PROGRESS"
	DONE        = "DONE"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func AddTask(description string) (*Task, error) {
	file, err := readDataFile()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("[!] Error while reading data from file: %w", err)
	}

	var lastId int

	if len(data) != 0 {

		var tasks []Task

		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return nil, fmt.Errorf("[!] Error while parsing json data: %w", err)
		}

		lastId = tasks[len(tasks)-1].Id

		newTask := &Task{
			Id:          lastId + 1,
			Description: description,
			Status:      TODO,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		encodedTasks, _ := json.Marshal(newTask)

		var whatToWrite []byte
		whatToWrite = append(whatToWrite, byte(','))
		whatToWrite = append(whatToWrite, encodedTasks...)
		whatToWrite = append(whatToWrite, byte(']'))

		file.WriteAt(whatToWrite, int64(len(data)-1))
		return newTask, nil

	} else {
		newTask := []Task{{
			Id:          lastId + 1,
			Description: description,
			Status:      TODO,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
		}
		encodedTasks, _ := json.Marshal(&newTask)
		file.Write(encodedTasks)

		return &newTask[0], nil
	}

}

func UpdateTask(id int, description string) (*Task, error) {
	file, err := readDataFile()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("[!] Error while reading data from file: %w", err)
	}

	tasks, err := parseTasks(data)
	if err != nil {
		return nil, err
	}

	var modifiedTask *Task

	for _, task := range tasks {
		if task.Id == id {
			task.Description = description
			task.UpdatedAt = time.Now()
			modifiedTask = task
			break
		}
	}

	encodedTasks, err := json.Marshal(tasks)
	if err != nil {
		return nil, fmt.Errorf("[!] Error while updating data: %w", err)
	}

	file, _ = os.Create(file.Name())
	file.Write(encodedTasks)

	return modifiedTask, nil
}

func DeleteTask(id int) error {
	file, err := readDataFile()

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("[!] Error while reading data from file: %w", err)
	}

	tasks, err := parseTasks(data)
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
	
	encodedTasks, _ := json.Marshal(tasks)
	file, _ = os.Create(file.Name())
	file.Write(encodedTasks)

	return nil
}
