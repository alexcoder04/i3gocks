package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ClickMessage struct {
	Name      string `json:"name"`
	Button    int    `json:"button"`
	Event     int    `json:"event"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	RelativeX int    `json:"relative_x"`
	RelativeY int    `json:"relative_y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Scale     int    `json:"scale"`
}

func ReadInput() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		// decode message
		line = strings.Trim(line, "[], \n")
		clickMsg := ClickMessage{}
		err = json.Unmarshal([]byte(line), &clickMsg)
		if err != nil {
			continue
		}

		// update clicked field and re-draw
		mu.Lock()
		for i := 0; i < len(config.Modules); i++ {
			if clickMsg.Name == config.Modules[i].Name {
				config.Modules[i] = UpdateModule(
					config.Modules[i],
					0,
					[]string{fmt.Sprintf("BLOCK_BUTTON=%d", clickMsg.Button)})
			}
		}
		draw(0)
		mu.Unlock()
	}
}
