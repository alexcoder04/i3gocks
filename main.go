package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(`{"version": 1, "click_events": true}`)
	fmt.Println("[")
	i := 1
	for {
		fmt.Printf(`[{"full_text": "HELLO %d", "color": "#0000ff", "background": "#ffff00"}],`+"\n", i)
		time.Sleep(1 * time.Second)
		i += 1
	}
}
