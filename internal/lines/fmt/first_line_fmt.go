package fmt

import "docker-color-output/pkg/color"

type FirstLineFmt struct{}

func NewFirstLineFmt() *FirstLineFmt {
	return &FirstLineFmt{}
}

func (f *FirstLineFmt) Format(_ string, v string) string {
	return color.LightBlue(v)
}
