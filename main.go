package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Module struct {
	Text            string   `json:"full_text"`
	ForegroundColor string   `json:"color"`
	BackgroundColor string   `json:"background"`
	Command         string   `json:"-"`
	Args            []string `json:"-"`
}

func UpdateModule(module Module) Module {
	cmd := exec.Command(module.Command, module.Args...)
	out, err := cmd.Output()
	if err != nil {
		module.Text = "error"
	} else {
		module.Text = strings.Replace(string(out), "\n", " ", -1)
	}
	return module
}

func main() {
	fmt.Println(`{"version": 1, "click_events": true}`)
	fmt.Println("[")

	modules := []Module{
		{"", "#0000ff", "#ffff00", "echo", []string{"kherson"}},
		{"", "#ffffff", "#000000", "date", []string{"+%d.%m.%Y - %R:%S"}}}

	i := 1
	for {
		listJson := ""
		for i := 0; i < len(modules); i++ {
			modules[i] = UpdateModule(modules[i])
			moduleJson, err := json.Marshal(modules[i])
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
