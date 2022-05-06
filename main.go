package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Module struct {
	Text            string   `json:"full_text" yaml:"-"`
	ForegroundColor string   `json:"color" yaml:"ForegroundColor"`
	BackgroundColor string   `json:"background" yaml:"BackgroundColor"`
	Icon            string   `json:"-" yaml:"Icon"`
	Command         string   `json:"-" yaml:"Command"`
	Args            []string `json:"-" yaml:"Args"`
	Interval        int      `json:"-" yaml:"Interval"`
}

type Config struct {
	Modules []Module `yaml:"Modules"`
}

func UpdateModule(module Module, counter int) Module {
	if counter%module.Interval != 0 {
		return module
	}

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

	counter := 0
	for {
		listJson := ""
		for i := 0; i < len(config.Modules); i++ {
			config.Modules[i] = UpdateModule(config.Modules[i], counter)
			moduleJson, err := json.Marshal(config.Modules[i])
			if err != nil {
				panic("error marshaling")
			}
			if listJson == "" {
				listJson = string(moduleJson)
				continue
			}
			listJson = listJson + "," + string(moduleJson)
		}

		fmt.Printf("[" + listJson + "],\n")
		time.Sleep(1 * time.Second)
		counter += 1
	}
}
