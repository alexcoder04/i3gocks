package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Module struct {
	Text            string   `json:"full_text" yaml:"Text"`
	ForegroundColor string   `json:"color" yaml:"ForegroundColor"`
	BackgroundColor string   `json:"background" yaml:"BackgroundColor"`
	Icon            string   `json:"-" yaml:"Icon"`
	Command         string   `json:"-" yaml:"Command"`
	Args            []string `json:"-" yaml:"Args"`
}

type Config struct {
	Modules []Module `yaml:"Modules"`
}

func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("cannot get config directory")
	}
	config := Config{}
	configFile := path.Join(configDir, "kherson", "config.yml")
	_, err = os.Stat(configFile)
	if err != nil {
		config.Modules = []Module{
			{"", "#0000ff", "#ffff00", " ", "echo", []string{"kherson"}},
			{"", "#ffffff", "#000000", " ", "date", []string{"+%d.%m.%Y - %R:%S"}}}
		return config
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("cannot read config file")
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("error parsing config")
	}
	return config
}

func UpdateModule(module Module) Module {
	cmd := exec.Command(module.Command, module.Args...)
	out, err := cmd.Output()
	if err != nil {
		module.Text = " error"
	} else {
		module.Text = module.Icon + strings.Replace(string(out), "\n", " ", -1)
	}
	return module
}

func main() {
	config := LoadConfig()

	fmt.Println(`{"version": 1, "click_events": true}`)
	fmt.Println("[")

	i := 1
	for {
		listJson := ""
		for i := 0; i < len(config.Modules); i++ {
			config.Modules[i] = UpdateModule(config.Modules[i])
			moduleJson, err := json.Marshal(config.Modules[i])
			if err != nil {
				panic("error marshaling")
			}
			if listJson == "" {
				listJson = "[" + string(moduleJson)
				continue
			}
			listJson = listJson + "," + string(moduleJson)
		}
		listJson = listJson + "],\n"

		fmt.Printf(listJson)
		time.Sleep(1 * time.Second)
		i += 1
	}
}
