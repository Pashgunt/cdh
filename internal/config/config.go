package config

type Config struct {
	Name    string           `yaml:"name" required:"true" default:"go-dir"`
	File    string           `yaml:"file" required:"true" default:".dirhistory"`
	Home    string           `yaml:"home" required:"true" default:"/"`
	Storage int              `yaml:"storage" required:"true" default:"100"`
	Error   map[string]Error `yaml:"error" required:"true"`
}

type Error struct {
	En string `json:"en"`
	Ru string `json:"ru"`
}
