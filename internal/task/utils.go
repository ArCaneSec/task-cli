package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func openDataFile() (*os.File, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("[!] An error occured when tried to open data file: %w", err)
	}

	exePath = filepath.Clean(exePath)
	dir := filepath.Dir(exePath)

	file, err := os.OpenFile(fmt.Sprintf("%s/data.json", dir), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("[!] Error while reading data file: %w", err)
	}

	return file, nil
}

func readDataFile() ([]byte, error) {
	file, err := openDataFile()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("[!] Error while reading data from file: %w", err)
	}

	return data, nil
}

func getTasks() ([]*Task, error) {
	var tasks []*Task

	data, err := readDataFile()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("[!] Error while parsing tasks: %w", err)
	}

	return tasks, nil
}

func updateDataFile(data []byte) error {
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
