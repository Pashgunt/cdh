package service

import (
	"bufio"
	"dir/internal/config"
	"dir/internal/helper"
	"os"
	"strings"
)

type HistoryContract interface {
	ReadHistory() ([]string, error)
	WriteHistory(history []string) error
	RemoveDirFromHistory(history []string, dir string) []string
	ClearHistory() error
}

type History struct {
	config config.Config
}

func NewHistory(config config.Config) *History {
	return &History{config: config}
}

func (h *History) ReadHistory() ([]string, error) {
	filePath := helper.GetHistoryFilePath(h.config)
	file, err := os.Open(filePath)

	if os.IsNotExist(err) {
		return []string{}, nil
	}

	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var history []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dir := strings.TrimSpace(scanner.Text())
		if dir != "" {
			history = append(history, dir)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return history, nil
}

func (h *History) WriteHistory(history []string) error {
	filePath := helper.GetHistoryFilePath(h.config)
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for _, dir := range history {
		_, err := file.WriteString(dir + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *History) RemoveDirFromHistory(history []string, dir string) []string {
	var newHistory []string

	for _, hDir := range history {
		if hDir == dir {
			continue
		}

		newHistory = append(newHistory, hDir)
	}

	return newHistory
}

func (h *History) ClearHistory() error {
	return h.WriteHistory([]string{})
}
