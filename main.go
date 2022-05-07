package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
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
}

type Config struct {
	Modules []Module `yaml:"Modules"`
}

var config Config
var mu sync.Mutex

func UpdateModule(module Module, counter int, env []string) Module {
	if counter%module.Interval != 0 {
		return module
	}

	cmd := exec.Command(module.Command, module.Args...)
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.Output()
	if err != nil {
		module.Text = " error"
	} else {
		lines := strings.Split(string(out), "\n")
		module.Text = fmt.Sprintf("%s%s%s",
			module.Pre,
			strings.Replace(lines[0], "\n", " ", -1),
			module.Post)
	}
	return module
}

func main() {
	config = LoadConfig()

	fmt.Println(`{"version": 1, "click_events": true}`)
	fmt.Println(`[`)

	go ReadInput()

	counter := 0
	listJson := ""
	for {
		listJson = ""
		for i := 0; i < len(config.Modules); i++ {
			config.Modules[i] = UpdateModule(config.Modules[i], counter, []string{})

			moduleJson, err := json.Marshal(config.Modules[i])
			if err != nil {
				moduleJson = []byte(`{"full_text":" error"}`)
			}
			if listJson == "" {
				listJson = string(moduleJson)
				continue
			}
			listJson = listJson + "," + string(moduleJson)
		}

		fmt.Printf("[%s],\n", listJson)
		time.Sleep(1 * time.Second)
		counter += 1
	}
}
