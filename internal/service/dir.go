package service

import (
	"fmt"
	"github.com/Pashgunt/cdh/internal/config"
)

type DirContract interface {
	AddDir(dir string) error
	ListDirs() error
	GetDir(index int) error
}

type Dir struct {
	config  config.Config
	history *History
}

func NewDir(config config.Config, history *History) *Dir {
	return &Dir{config: config, history: history}
}

func (d *Dir) AddDir(dir string) error {
	history, err := d.history.ReadHistory()

	if err != nil {
		return err
	}

	history = append(history, dir)

	if len(history) > d.config.Storage {
		history = history[:d.config.Storage]
	}

	return d.history.WriteHistory(history)
}

func (d *Dir) ListDirs() error {
	history, err := d.history.ReadHistory()

	if err != nil {
		return err
	}

	for i, dir := range history {
		fmt.Printf("%d: %s\n", i+1, dir)
	}

	return nil
}

func (d *Dir) GetDir(index int) error {
	history, err := d.history.ReadHistory()

	if err != nil {
		return err
	}

	if index < 0 || index >= len(history) {
		return fmt.Errorf("индекс вне диапазона")
	}

	fmt.Print(history[index])

	return nil
}
