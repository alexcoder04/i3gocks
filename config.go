package main

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func DefaultConfig(msg string) Config {
	config := Config{}
	config.Modules = []Module{
		{"", "#0000ff", "#ffff00", " ", "echo", []string{msg}, 60},
		{"", "#ffffff", "#000000", " ", "date", []string{"+%d.%m.%Y - %R:%S"}, 1},
		{"", "#880088", "#ccccee", " ", "uname", []string{"-r"}, 60}}
	return config
}

func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return DefaultConfig("cannot get config dir")
	}

	configFile := path.Join(configDir, "kherson", "config.yml")
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultConfig("kherson (default config)")
		}
		return DefaultConfig("cannot stat config")
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return DefaultConfig("error reading config")
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return DefaultConfig("error parsing config")
	}
	return config
}
