package fmt

import "docker-color-output/pkg/color"

type FirstLineFmt struct{}

func NewFirstLineFmt() *FirstLineFmt {
	return &FirstLineFmt{}
}

func (f *FirstLineFmt) Format(_ map[string]string, col string) string {
	return color.LightBlue(col)
}
