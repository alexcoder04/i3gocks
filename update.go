package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func UpdateModuleByName(name string, counter int, env []string) {
	for i := 0; i < len(config.Modules); i++ {
		if name == config.Modules[i].Name {
			UpdateModule(i, counter, env)
		}
	}
}

func UpdateModule(mod int, counter int, env []string) {
	// don't update if interval didn't pass
	if counter%config.Modules[mod].Interval != 0 {
		return
	}

	cmd := exec.Command(config.Modules[mod].Command, config.Modules[mod].Args...)
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.Output()

	if err != nil {
		config.Modules[mod].Text = " error"
	} else {
		lines := strings.Split(string(out), "\n")
		for i := 0; i <= 3; i++ {
			if len(lines) < i+1 {
				break
			}
			switch i {
			// first line is text
			case 0:
				config.Modules[mod].Text = fmt.Sprintf("%s%s%s",
					config.Modules[mod].Pre,
					strings.Replace(lines[i], "\n", " ", -1),
					config.Modules[mod].Post)
			// third line is ForegroundColor
			case 2:
				config.Modules[mod].ForegroundColor = lines[i]
			// fourth line is BackgroundColor
			case 3:
				config.Modules[mod].BackgroundColor = lines[i]
			}
		}
	}
}
