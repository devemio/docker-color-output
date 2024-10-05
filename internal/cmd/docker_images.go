package cmd

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/pkg/color"
	"github.com/devemio/docker-color-output/pkg/util/number"
)

const (
	DockerImagesRepository = "REPOSITORY"
	DockerImagesTag        = "TAG"
	DockerImagesDigest     = "DIGEST"
	DockerImagesImageID    = "IMAGE ID"
	DockerImagesCreated    = "CREATED"
	DockerImagesCreatedAt  = "CREATED AT"
	DockerImagesSize       = "SIZE"
)

type DockerImages struct{}

func (c *DockerImages) Columns() []string {
	return []string{
		DockerImagesImageID,    //
		DockerImagesRepository, //
		DockerImagesTag,        //
		DockerImagesDigest,     // opt
		DockerImagesCreated,    //
		DockerImagesCreatedAt,  // opt
		DockerImagesSize,       //
	}
}

func (c *DockerImages) Format(row layout.Row, col layout.Column) string {
	x := string(row[col])

	switch col {
	case DockerImagesRepository:
		return c.Repository(x)
	case DockerImagesTag:
		return c.Tag(x)
	case DockerImagesImageID:
		return c.ImageID(x)
	case DockerImagesCreated:
		return c.Created(x)
	case DockerImagesSize:
		return c.Size(x)
	default:
		return x
	}
}

func (*DockerImages) Repository(x string) string {
	if strings.Contains(x, "/") {
		return color.DarkGray(x)
	}

	return color.White(x)
}

func (*DockerImages) Tag(x string) string {
	if x == "latest" {
		return color.LightGreen(x)
	}

	return x
}

func (*DockerImages) ImageID(x string) string {
	return color.DarkGray(x)
}

func (*DockerImages) Created(x string) string {
	if strings.Contains(x, "hour") {
		return color.Green(x)
	}

	if strings.Contains(x, "days") {
		return color.Green(x)
	}

	if strings.Contains(x, "weeks") {
		return color.Green(x)
	}

	if strings.Contains(x, "months") {
		return color.Brown(x)
	}

	if strings.Contains(x, "years") {
		return color.Red(x)
	}

	return x
}

func (*DockerImages) Size(x string) string {
	if strings.Contains(x, "GB") {
		return color.Red(x)
	}

	if strings.Contains(x, "MB") && number.ParseFloat(x) >= 500 {
		return color.Brown(x)
	}

	return x
}
