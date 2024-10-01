package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func readDataFile() (*os.File, error) {
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

func parseTasks(data []byte) ([]*Task, error) {
	var tasks []*Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("[!] Error while parsing tasks: %w", err)
	}

	return tasks, nil
}