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
			config.Modules[i] = UpdateModule(
				config.Modules[i],
				counter,
				env)
		}
	}
}

func UpdateModule(module Module, counter int, env []string) Module {
	// don't update if interval didn't pass
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
			// first line is text
			case 0:
				module.Text = fmt.Sprintf("%s%s%s",
					module.Pre,
					strings.Replace(lines[i], "\n", " ", -1),
					module.Post)
			// third line is ForegroundColor
			case 2:
				module.ForegroundColor = lines[i]
			// fourth line is BackgroundColor
			case 3:
				module.BackgroundColor = lines[i]
			}
		}
	}
	return module
}
