package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func taskFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func ReadTasksFromFile() ([]Task, error) {
	filePath := taskFilePath()
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		return []Task{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	task := []Task{}
	err = json.NewDecoder(file).Decode(&task)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return task, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath := taskFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}
