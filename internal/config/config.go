package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/devemio/docker-color-output/internal/app"
)

type Config struct {
	Colors Colors `json:"colors"`
}

type Colors struct {
	Reset       string `json:"reset"`
	Black       string `json:"black"`
	DarkGray    string `json:"darkGray"`
	Red         string `json:"red"`
	LightRed    string `json:"lightRed"`
	Green       string `json:"green"`
	LightGreen  string `json:"lightGreen"`
	Brown       string `json:"brown"`
	Yellow      string `json:"yellow"`
	Blue        string `json:"blue"`
	LightBlue   string `json:"lightBlue"`
	Purple      string `json:"purple"`
	LightPurple string `json:"lightPurple"`
	Cyan        string `json:"cyan"`
	LightCyan   string `json:"lightCyan"`
	LightGray   string `json:"lightGray"`
	White       string `json:"white"`
}

func Get() (Config, error) {
	cfg := createDefault()

	flag.Usage = func() {
		app.Usage(nil)
	}

	cfgPath := flag.String("c", "", "Path to configuration file")
	flag.Parse()

	if *cfgPath != "" {
		data, err := os.ReadFile(*cfgPath)
		if err != nil {
			return Config{}, fmt.Errorf("read: %w", err)
		}

		if err = json.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("unmarshal: %w", err)
		}
	}

	return cfg, nil
}

func createDefault() Config {
	return Config{
		Colors: Colors{
			Reset:       "\u001B[0m",
			Black:       "\u001B[0;30m",
			DarkGray:    "\u001B[1;30m",
			Red:         "\u001B[0;31m",
			LightRed:    "\u001B[1;31m",
			Green:       "\u001B[0;32m",
			LightGreen:  "\u001B[1;32m",
			Brown:       "\u001B[0;33m",
			Yellow:      "\u001B[1;33m",
			Blue:        "\u001B[0;34m",
			LightBlue:   "\u001B[1;34m",
			Purple:      "\u001B[0;35m",
			LightPurple: "\u001B[1;35m",
			Cyan:        "\u001B[0;36m",
			LightCyan:   "\u001B[1;36m",
			LightGray:   "\u001B[0;37m",
			White:       "\u001B[1;37m",
		},
	}
}
