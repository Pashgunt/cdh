package main

import (
	_ "embed"
	"github.com/Pashgunt/cdh/internal/config"
	"github.com/Pashgunt/cdh/internal/helper"
	"github.com/Pashgunt/cdh/internal/service"
	"gopkg.in/yaml.v2"
	"os"
)

//go:embed config.yaml
var embeddedConfig []byte

func main() {
	var settings config.Config
	_ = yaml.Unmarshal(embeddedConfig, &settings)
	history := service.NewHistory(settings)
	dir := service.NewDir(settings, history)

	if len(os.Args) < 2 {
		helper.PrintUsage(settings)

		return
	}

	if executor, isset := service.GetCommandExecution(settings, dir, history)[os.Args[1]]; isset {
		executor()

		return
	}

	helper.PrintUsage(settings)
}
