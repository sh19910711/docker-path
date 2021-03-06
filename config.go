package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Command struct {
	Image   string   `image`
	Command string   `command`
	Git     bool     `git`
	Env     []string `env`
	ExecEnv []string `exec_env`
}

type Config map[string]Command

func Parse(text string) Config {
	bytes := []byte(text)
	config := make(Config)

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
