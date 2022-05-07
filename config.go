package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type Module struct {
	Name            string   `json:"name" yaml:"Name"`
	Text            string   `json:"full_text" yaml:"-"`
	ForegroundColor string   `json:"color" yaml:"ForegroundColor"`
	BackgroundColor string   `json:"background" yaml:"BackgroundColor"`
	Pre             string   `json:"-" yaml:"Pre"`
	Post            string   `json:"-" yaml:"Post"`
	Command         string   `json:"-" yaml:"Command"`
	Args            []string `json:"-" yaml:"Args"`
	Interval        int      `json:"-" yaml:"Interval"`
	Markup          string   `json:"markup" yaml:"Markup"`
}

type Config struct {
	Modules []Module          `yaml:"Modules"`
	Colors  map[string]string `yaml:"Colors"`
}

func DefaultConfig(msg string) Config {
	config := Config{}
	config.Modules = []Module{
		{"msg", "", "#0000ff", "#ffff00", " ", "", "echo", []string{msg}, 60, "none"},
		{"time", "", "#ffffff", "#000000", " ", "", "date", []string{"+%d.%m.%Y - %R:%S"}, 1, "none"},
		{"kernel", "", "#880088", "#ccccee", " ", "", "uname", []string{"-r"}, 60, "none"}}
	return config
}

func LoadColors() map[string]string {
	// default
	colors := map[string]string{
		"BLACK":        "#282828",
		"BLUE":         "#458588",
		"CYAN":         "#689d6a",
		"DARK_BLUE":    "#458588",
		"DARK_GREY":    "#6f6357",
		"GREEN":        "#98971a",
		"LIGHT_BLUE":   "#83a598",
		"LIGHT_CYAN":   "#8ec07c",
		"LIGHT_GREEN":  "#b8bb26",
		"LIGHT_GREY":   "#a89984",
		"LIGHT_PURPLE": "#d3869b",
		"LIGHT_RED":    "#fb4934",
		"LIGHT_YELLOW": "#fabd2f",
		"MAGENTA":      "#b16286",
		"PURPLE":       "#b16286",
		"RED":          "#cc241d",
		"WHITE":        "#ebdbb2",
		"YELLOW":       "#d79921"}

	// load from env
	for _, v := range os.Environ() {
		keyValue := strings.Split(v, "=")
		if len(keyValue[0]) < 7 {
			continue
		}
		if keyValue[0][:6] == "COLOR_" {
			colors[keyValue[0][6:]] = keyValue[1]
		}
	}
	return colors
}

func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return DefaultConfig("cannot get config dir")
	}

	configFile := path.Join(configDir, "kherson", "config.yml")
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultConfig("kherson (default config)")
		}
		return DefaultConfig("cannot stat config")
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return DefaultConfig("error reading config")
	}

	config := Config{}
	config.Colors = LoadColors()
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return DefaultConfig("error parsing config")
	}
	for i := 0; i < len(config.Modules); i++ {
		if config.Modules[i].ForegroundColor == "" {
			config.Modules[i].ForegroundColor = config.Colors["WHITE"]
		}
		if config.Modules[i].ForegroundColor[0] == '*' {
			config.Modules[i].ForegroundColor = config.Colors[strings.ToUpper(config.Modules[i].ForegroundColor[1:])]
		}
		if config.Modules[i].BackgroundColor != "" && config.Modules[i].BackgroundColor[0] == '*' {
			config.Modules[i].BackgroundColor = config.Colors[strings.ToUpper(config.Modules[i].BackgroundColor[1:])]
		}
	}
	return config
}
