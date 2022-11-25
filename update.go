package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetPowerlinePrefix(mod int) string {
	if !config.Options.PowerlineTheme {
		return ""
	}
	// foreground (symbol itself) is our background
	fg := config.Modules[mod].BackgroundColor
	// background is the background of previous shown module
	var bg string
	if mod == 0 {
		bg = config.Colors["BLACK"]
	} else {
		i := 0
		for {
			i += 1
			if i < 0 {
				bg = config.Colors["BLACK"]
				break
			}
			if config.Modules[mod-i].Text != "" {
				bg = config.Modules[mod-i].BackgroundColor
				break
			}
		}
	}
	return fmt.Sprintf("<span foreground='%s' background='%s'>%s</span>", fg, bg, config.Options.PowerlineSeparator)
}

func UpdateModuleByName(name string, counter int, env []string) {
	for i := 0; i < len(config.Modules); i++ {
		if name == config.Modules[i].Name {
			UpdateModule(i, counter, env)
		}
	}
}

func ExecuteCommand(command string, args []string, env []string) []string {
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.Output()

	if err != nil {
		return []string{" error"}
	}

	return strings.Split(string(out), "\n")
}

func UpdateModule(mod int, counter int, env []string) {
	// don't update if interval didn't pass
	if counter%config.Modules[mod].Interval != 0 {
		return
	}

	var lines []string
	if config.Modules[mod].Command[0] == '*' {
		lines = BuiltIns[config.Modules[mod].Command[1:]](config.Modules[mod].Args)
	} else {
		lines = ExecuteCommand(config.Modules[mod].Command, config.Modules[mod].Args, env)
	}

	for i := 0; i <= len(lines)-1; i++ {
		switch i {
		// first line is text
		case 0:
			if lines[i] == "" {
				break
			}
			config.Modules[mod].Text = fmt.Sprintf("%s%s%s%s",
				GetPowerlinePrefix(mod),
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
