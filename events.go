package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

const SIGRTMIN = 34

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
		UpdateModuleByName(
			clickMsg.Name,
			0,
			[]string{fmt.Sprintf("BLOCK_BUTTON=%d", clickMsg.Button)})
		draw()
		mu.Unlock()
	}
}

func ListenFor(signalNumber int, blockName string) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(SIGRTMIN+signalNumber))
	select {
	case <-channel:
		mu.Lock()
		UpdateModuleByName(blockName, 0, []string{})
		draw()
		mu.Unlock()
		ListenFor(signalNumber, blockName)
	}
}

func ListenToReloadConfig() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.Signal(syscall.SIGUSR1))
	select {
	case <-channel:
		mu.Lock()
		fmt.Printf(
			`[{"full_text": "reloading config...", "color": "%s"}],`,
			config.Colors["WHITE"])
		config = LoadConfig()
		for i := 0; i < len(config.Modules); i++ {
			UpdateModule(i, 0, []string{})
		}
		draw()
		mu.Unlock()

		ListenToReloadConfig()
	}
}
