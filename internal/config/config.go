package config

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/devemio/docker-color-output/internal/app"
	"github.com/devemio/docker-color-output/internal/rules"
	"gopkg.in/yaml.v3"
)

type Config struct {
	SilentMode bool         `yaml:"-"`
	Colors     Colors       `yaml:"colors"`
	Layout     Layout       `yaml:"layout"`
	Rules      rules.Config `yaml:"rules"`
}

type Colors struct {
	Reset       string `yaml:"reset"`
	Black       string `yaml:"black"`
	DarkGray    string `yaml:"darkGray"`
	Red         string `yaml:"red"`
	LightRed    string `yaml:"lightRed"`
	Green       string `yaml:"green"`
	LightGreen  string `yaml:"lightGreen"`
	Brown       string `yaml:"brown"`
	Yellow      string `yaml:"yellow"`
	Blue        string `yaml:"blue"`
	LightBlue   string `yaml:"lightBlue"`
	Purple      string `yaml:"purple"`
	LightPurple string `yaml:"lightPurple"`
	Cyan        string `yaml:"cyan"`
	LightCyan   string `yaml:"lightCyan"`
	LightGray   string `yaml:"lightGray"`
	White       string `yaml:"white"`
}

type Layout struct {
	HeaderColor string `yaml:"headerColor"`
}

//go:embed default.yaml
var defaultConfig []byte

func Get() (Config, error) {
	cfg := createDefault()

	flag.Usage = func() {
		app.Usage(nil)
	}

	cfgPath := flag.String("c", "", "Path to configuration file")
	silentMode := flag.Bool("s", false, "Silent mode (suppress errors)")
	flag.Parse()

	if *cfgPath != "" {
		data, err := os.ReadFile(*cfgPath)
		if err != nil {
			return Config{}, fmt.Errorf("read: %w", err)
		}

		var override Config
		if err = yaml.Unmarshal(data, &override); err != nil {
			return Config{}, fmt.Errorf("unmarshal: %w", err)
		}

		cfg = merge(cfg, override)
	}

	if *silentMode {
		cfg.SilentMode = true
	}

	return cfg, nil
}

func Default() Config {
	return createDefault()
}

func createDefault() Config {
	var cfg Config
	if err := yaml.Unmarshal(defaultConfig, &cfg); err != nil {
		return Config{}
	}

	return cfg
}

func merge(base Config, override Config) Config {
	base.Colors = mergeColors(base.Colors, override.Colors)
	base.Layout = mergeLayout(base.Layout, override.Layout)
	base.Rules = mergeRules(base.Rules, override.Rules)

	return base
}

func mergeColors(base Colors, override Colors) Colors {
	if override.Reset != "" {
		base.Reset = override.Reset
	}
	if override.Black != "" {
		base.Black = override.Black
	}
	if override.DarkGray != "" {
		base.DarkGray = override.DarkGray
	}
	if override.Red != "" {
		base.Red = override.Red
	}
	if override.LightRed != "" {
		base.LightRed = override.LightRed
	}
	if override.Green != "" {
		base.Green = override.Green
	}
	if override.LightGreen != "" {
		base.LightGreen = override.LightGreen
	}
	if override.Brown != "" {
		base.Brown = override.Brown
	}
	if override.Yellow != "" {
		base.Yellow = override.Yellow
	}
	if override.Blue != "" {
		base.Blue = override.Blue
	}
	if override.LightBlue != "" {
		base.LightBlue = override.LightBlue
	}
	if override.Purple != "" {
		base.Purple = override.Purple
	}
	if override.LightPurple != "" {
		base.LightPurple = override.LightPurple
	}
	if override.Cyan != "" {
		base.Cyan = override.Cyan
	}
	if override.LightCyan != "" {
		base.LightCyan = override.LightCyan
	}
	if override.LightGray != "" {
		base.LightGray = override.LightGray
	}
	if override.White != "" {
		base.White = override.White
	}

	return base
}

func mergeLayout(base Layout, override Layout) Layout {
	if override.HeaderColor != "" {
		base.HeaderColor = override.HeaderColor
	}

	return base
}

func mergeRules(base rules.Config, override rules.Config) rules.Config {
	base.DockerPs = mergeCommand(base.DockerPs, override.DockerPs)
	base.DockerImages = mergeCommand(base.DockerImages, override.DockerImages)
	base.DockerComposePs = mergeCommand(base.DockerComposePs, override.DockerComposePs)
	base.DockerStats = mergeCommand(base.DockerStats, override.DockerStats)

	return base
}

func mergeCommand(base rules.CommandConfig, override rules.CommandConfig) rules.CommandConfig {
	if base.Columns == nil {
		base.Columns = make(map[string]rules.ColumnRule)
	}
	for name, rule := range override.Columns {
		base.Columns[name] = rule
	}

	return base
}
