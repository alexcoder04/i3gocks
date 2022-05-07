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
		for i := 0; i <= 3; i++ {
			if len(lines) < i+1 {
				break
			}
			switch i {
			case 0:
				module.Text = fmt.Sprintf("%s%s%s",
					module.Pre,
					strings.Replace(lines[i], "\n", " ", -1),
					module.Post)
			case 2:
				module.ForegroundColor = lines[i]
			case 3:
				module.BackgroundColor = lines[i]
			}
		}
	}
	return module
}

func draw(counter int) {
	listJson := ""
	for i := 0; i < len(config.Modules); i++ {
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
}

func main() {
	config = LoadConfig()

	fmt.Println(`{"version": 1, "click_events": true}`)
	fmt.Println(`[`)
	fmt.Printf(
		`[{"full_text": "loading status line...", "color": "%s"}],`,
		config.Colors["WHITE"])

	go ReadInput()

	counter := 0
	for {
		for i := 0; i < len(config.Modules); i++ {
			config.Modules[i] = UpdateModule(config.Modules[i], counter, []string{})
		}
		draw(counter)
		time.Sleep(1 * time.Second)
		counter += 1
	}
}
