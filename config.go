package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	configLocation = flag.String("config", "", "config file")
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
	Signal          int      `json:"-" yaml:"Signal"`
}

type ConfigOptions struct {
	PowerlineTheme bool `yaml:"PowerlineTheme"`
}

type Config struct {
	Modules []Module          `yaml:"Modules"`
	Colors  map[string]string `yaml:"Colors"`
	Options ConfigOptions     `yaml:"Options"`
}

func DefaultConfig(msg string) Config {
	config := Config{}
	config.Modules = []Module{
		{"msg", "", "#0000ff", "#ffff00", " ", "",
			"echo", []string{msg}, 60, "none", 0},
		{"time", "", "#ffffff", "#000000", " ", "",
			"date", []string{"+%d.%m.%Y - %R:%S"}, 1, "none", 0},
		{"kernel", "", "#880088", "#ccccee", " ", "",
			"uname", []string{"-r"}, 60, "none", 0}}
	return config
}

func LoadColors() map[string]string {
	// default (gruvbox)
	colors := map[string]string{
		"BLACK":      "#282828",
		"BLUE":       "#458588",
		"CYAN":       "#689d6a",
		"DARK_GREY":  "#6f6357",
		"GREEN":      "#98971a",
		"LIGHT_GREY": "#a89984",
		"MAGENTA":    "#b16286",
		"RED":        "#cc241d",
		"WHITE":      "#ebdbb2",
		"YELLOW":     "#d79921"}

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
	flag.Parse()

	var configFile string

	if *configLocation != "" {
		configFile = *configLocation
	} else {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return DefaultConfig("cannot get config dir")
		}
		configFile = path.Join(configDir, "kherson", "config.yml")
	}

	_, err := os.Stat(configFile)
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
		// default interval
		if config.Modules[i].Interval == 0 {
			config.Modules[i].Interval = 1
		}
		// default foreground color
		if config.Modules[i].ForegroundColor == "" {
			config.Modules[i].ForegroundColor = config.Colors["WHITE"]
		}
		// default background color
		if config.Modules[i].BackgroundColor == "" {
			config.Modules[i].BackgroundColor = config.Colors["BLACK"]
		}
		// foreground color reference
		if config.Modules[i].ForegroundColor[0] == '*' {
			config.Modules[i].ForegroundColor = config.Colors[strings.ToUpper(config.Modules[i].ForegroundColor[1:])]
		}
		// background color reference
		if config.Modules[i].BackgroundColor != "" && config.Modules[i].BackgroundColor[0] == '*' {
			config.Modules[i].BackgroundColor = config.Colors[strings.ToUpper(config.Modules[i].BackgroundColor[1:])]
		}
		// enable pango on every module in powerline theme
		if config.Options.PowerlineTheme {
			config.Modules[i].Markup = "pango"
		}
	}
	return config
}
