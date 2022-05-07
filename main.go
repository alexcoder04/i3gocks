package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var config Config
var mu sync.Mutex

func draw() {
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
	for _, m := range config.Modules {
		if m.Signal >= 1 && m.Signal <= 15 {
			go ListenFor(34+m.Signal, m.Name)
		}
	}

	counter := 0
	for {
		for i := 0; i < len(config.Modules); i++ {
			config.Modules[i] = UpdateModule(config.Modules[i], counter, []string{})
		}

		draw()
		time.Sleep(1 * time.Second)
		counter += 1
	}
}
