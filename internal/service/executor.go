package service

import (
	"dir/internal/config"
	"dir/internal/helper"
	"fmt"
	"os"
	"strconv"
)

func GetCommandExecution(
	settings config.Config,
	dir *Dir,
	history *History,
) map[string]func() {
	return map[string]func(){
		helper.CommandAdd: func() {
			var path string

			if len(os.Args) >= 3 {
				path = os.Args[2]
			} else {
				var err error
				path, err = os.Getwd()

				if err != nil {
					fmt.Println(settings.Error["currentDir"].En, err)

					return
				}
			}

			if err := dir.AddDir(path); err != nil {
				fmt.Println(settings.Error["addDirToHistory"].En, err)
			}
		},
		helper.CommandList: func() {
			if err := dir.ListDirs(); err != nil {
				fmt.Println(settings.Error["historyList"].En, err)
			}
		},
		helper.CommandGet: func() {
			if len(os.Args) < 3 {
				fmt.Println(settings.Error["dirIndex"].En)

				return
			}

			index, err := strconv.Atoi(os.Args[2])

			if err != nil {
				fmt.Println(settings.Error["incorrectDirIndex"].En, os.Args[2])

				return
			}

			if err := dir.GetDir(index - 1); err != nil {
				fmt.Println(settings.Error["getDir"].En, err)
			}
		},
		helper.CommandClear: func() {
			if err := history.ClearHistory(); err != nil {
				fmt.Println(settings.Error["clearHistory"].En, err)
			}
		},
	}
}
