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

func UpdateModule(i int, counter int, env []string) {
	// don't update if interval didn't pass
	if counter%config.Modules[i].Interval != 0 {
		return
	}

	cmd := exec.Command(config.Modules[i].Command, config.Modules[i].Args...)
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.Output()

	if err != nil {
		config.Modules[i].Text = " error"
	} else {
		lines := strings.Split(string(out), "\n")
		for i := 0; i <= 3; i++ {
			if len(lines) < i+1 {
				break
			}
			switch i {
			// first line is text
			case 0:
				config.Modules[i].Text = fmt.Sprintf("%s%s%s",
					config.Modules[i].Pre,
					strings.Replace(lines[i], "\n", " ", -1),
					config.Modules[i].Post)
			// third line is ForegroundColor
			case 2:
				config.Modules[i].ForegroundColor = lines[i]
			// fourth line is BackgroundColor
			case 3:
				config.Modules[i].BackgroundColor = lines[i]
			}
		}
	}
}
