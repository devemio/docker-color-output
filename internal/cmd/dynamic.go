package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
)

type DynamicCommand struct {
	rules rules.CommandConfig
}

func (c *DynamicCommand) Format(row layout.Row, col layout.Column) string {
	x := string(row[col])

	if rule, ok := c.rules.Rules[string(col)]; ok {
		return rule.Pipeline.Apply(x, rules.Context{Row: row, Column: col})
	}

	return x
}
