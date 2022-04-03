package fmt

import (
	"strings"

	"docker-color-output/pkg/color"
	"docker-color-output/pkg/utils/numbers"
)

type DockerImagesLineFmt struct{}

func NewDockerImagesLineFmt() *DockerImagesLineFmt {
	return &DockerImagesLineFmt{}
}

func (*DockerImagesLineFmt) Repository(v string) string {
	if strings.Contains(v, "/") {
		return color.DarkGray(v)
	}
	return color.White(v)
}

func (*DockerImagesLineFmt) Tag(v string) string {
	if v == "latest" {
		return color.LightGreen(v)
	}
	return v
}

func (*DockerImagesLineFmt) ImageId(v string) string {
	return color.DarkGray(v)
}

func (*DockerImagesLineFmt) Created(v string) string {
	if strings.Contains(v, "hour") {
		return color.Green(v)
	}
	if strings.Contains(v, "days") {
		return color.Green(v)
	}
	if strings.Contains(v, "weeks") {
		return color.Green(v)
	}
	if strings.Contains(v, "months") {
		return color.Brown(v)
	}
	if strings.Contains(v, "years") {
		return color.Red(v)
	}
	return v
}

func (*DockerImagesLineFmt) Size(v string) string {
	if strings.Contains(v, "GB") {
		return color.Red(v)
	}
	if strings.Contains(v, "MB") && numbers.ParseFloat(v) >= 500 {
		return color.Brown(v)
	}
	return v
}

func (f *DockerImagesLineFmt) Format(col string, v string) string {
	switch col {
	case "REPOSITORY":
		return f.Repository(v)
	case "TAG":
		return f.Tag(v)
	case "IMAGE ID":
		return f.ImageId(v)
	case "CREATED":
		return f.Created(v)
	case "SIZE":
		return f.Size(v)
	default:
		return v
	}
}
