package rules

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/pkg/color"
	"github.com/devemio/docker-color-output/pkg/util/number"
)

type Config struct {
	DockerPs        CommandConfig `yaml:"dockerPs"`
	DockerImages    CommandConfig `yaml:"dockerImages"`
	DockerComposePs CommandConfig `yaml:"dockerComposePs"`
	DockerStats     CommandConfig `yaml:"dockerStats"`
}

type CommandConfig struct {
	Columns map[string]ColumnRule `yaml:"columns"`
}

type Pipeline []Step

type ColumnRule struct {
	Pipeline Pipeline `yaml:"pipeline"`
}

type Step struct {
	Type      string      `yaml:"type"`
	Color     string      `yaml:"color,omitempty"`
	When      []MatchCase `yaml:"when,omitempty"`
	Default   *Action     `yaml:"default,omitempty"`
	Delimiter string      `yaml:"delimiter,omitempty"`
	Parts     []SplitPart `yaml:"parts,omitempty"`
	Fallback  *Pipeline   `yaml:"fallback,omitempty"`
	Joiner    *Joiner     `yaml:"joiner,omitempty"`
	Item      *Pipeline   `yaml:"item,omitempty"`
}

type MatchCase struct {
	Column    string   `yaml:"column,omitempty"`
	Contains  string   `yaml:"contains,omitempty"`
	Equals    string   `yaml:"equals,omitempty"`
	NumberGte *float64 `yaml:"numberGte,omitempty"`
	NumberGt  *float64 `yaml:"numberGt,omitempty"`
	NumberLte *float64 `yaml:"numberLte,omitempty"`
	NumberLt  *float64 `yaml:"numberLt,omitempty"`
	Color     string   `yaml:"color,omitempty"`
}

type Action struct {
	Color string `yaml:"color,omitempty"`
}

type SplitPart struct {
	Prefix   string   `yaml:"prefix,omitempty"`
	Suffix   string   `yaml:"suffix,omitempty"`
	Pipeline Pipeline `yaml:"pipeline,omitempty"`
}

type Joiner struct {
	Value string `yaml:"value,omitempty"`
	Color string `yaml:"color,omitempty"`
}

type Context struct {
	Row    layout.Row
	Column layout.Column
}

func (p Pipeline) Apply(value string, ctx Context) string {
	res := value
	for _, step := range p {
		res = step.Apply(res, ctx)
	}

	return res
}

func (s Step) Apply(value string, ctx Context) string {
	switch s.Type {
	case "color":
		return Colorize(s.Color, value)
	case "match":
		for _, when := range s.When {
			if when.Matches(value, ctx) {
				return Colorize(when.Color, value)
			}
		}
		if s.Default != nil {
			return Colorize(s.Default.Color, value)
		}
		return value
	case "split":
		if s.Delimiter == "" {
			return value
		}
		parts := strings.Split(value, s.Delimiter)
		if len(parts) != len(s.Parts) {
			if s.Fallback != nil {
				return s.Fallback.Apply(value, ctx)
			}
			return value
		}
		resParts := make([]string, len(parts))
		for i, part := range parts {
			partValue := s.Parts[i].Prefix + part + s.Parts[i].Suffix
			res := s.Parts[i].Pipeline.Apply(partValue, ctx)
			resParts[i] = res
		}
		joiner := s.Delimiter
		if s.Joiner != nil {
			joiner = Colorize(s.Joiner.Color, s.Joiner.Value)
		}
		return strings.Join(resParts, joiner)
	case "list":
		if s.Delimiter == "" || s.Item == nil {
			return value
		}
		items := strings.Split(value, s.Delimiter)
		for i, item := range items {
			items[i] = s.Item.Apply(item, ctx)
		}
		joiner := s.Delimiter
		if s.Joiner != nil {
			joiner = Colorize(s.Joiner.Color, s.Joiner.Value)
		}
		return strings.Join(items, joiner)
	default:
		return value
	}
}

func (m MatchCase) Matches(value string, ctx Context) bool {
	if m.Column != "" {
		colValue, ok := ctx.Row[layout.Column(m.Column)]
		if !ok {
			return false
		}
		value = string(colValue)
	}

	if m.Contains != "" && !strings.Contains(value, m.Contains) {
		return false
	}
	if m.Equals != "" && value != m.Equals {
		return false
	}
	if m.NumberGte != nil && number.ParseFloat(value) < *m.NumberGte {
		return false
	}
	if m.NumberGt != nil && number.ParseFloat(value) <= *m.NumberGt {
		return false
	}
	if m.NumberLte != nil && number.ParseFloat(value) > *m.NumberLte {
		return false
	}
	if m.NumberLt != nil && number.ParseFloat(value) >= *m.NumberLt {
		return false
	}

	return true
}

func Colorize(name string, value string) string {
	if name == "" {
		return value
	}

	if fn, ok := colorFuncs[normalizeColorName(name)]; ok {
		return fn(value)
	}

	return value
}

func normalizeColorName(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if r >= 'a' && r <= 'z' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var colorFuncs = map[string]func(string) string{ //nolint:gochecknoglobals
	"black":       color.Black,
	"darkgray":    color.DarkGray,
	"red":         color.Red,
	"lightred":    color.LightRed,
	"green":       color.Green,
	"lightgreen":  color.LightGreen,
	"brown":       color.Brown,
	"yellow":      color.Yellow,
	"blue":        color.Blue,
	"lightblue":   color.LightBlue,
	"purple":      color.Purple,
	"lightpurple": color.LightPurple,
	"cyan":        color.Cyan,
	"lightcyan":   color.LightCyan,
	"lightgray":   color.LightGray,
	"white":       color.White,
}
