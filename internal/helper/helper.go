package helper

import (
	"dir/internal/config"
	"fmt"
	"os"
	"path/filepath"
)

func GetHistoryFilePath(settings config.Config) string {
	return filepath.Join(os.Getenv("HOME"), settings.File)
}

func PrintUsage(config config.Config) {
	fmt.Println("Usage:")

	for command, description := range GetCommandDescription() {
		fmt.Printf("%s %s %s\n", config.Name, command, description)
	}
}
