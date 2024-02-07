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
	v := string(row[col])

	switch col {
	case DockerImagesRepository:
		return c.Repository(v)
	case DockerImagesTag:
		return c.Tag(v)
	case DockerImagesImageID:
		return c.ImageID(v)
	case DockerImagesCreated:
		return c.Created(v)
	case DockerImagesSize:
		return c.Size(v)
	default:
		return v
	}
}

func (*DockerImages) Repository(v string) string {
	if strings.Contains(v, "/") {
		return color.DarkGray(v)
	}

	return color.White(v)
}

func (*DockerImages) Tag(v string) string {
	if v == "latest" {
		return color.LightGreen(v)
	}

	return v
}

func (*DockerImages) ImageID(v string) string {
	return color.DarkGray(v)
}

func (*DockerImages) Created(v string) string {
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

func (*DockerImages) Size(v string) string {
	if strings.Contains(v, "GB") {
		return color.Red(v)
	}

	if strings.Contains(v, "MB") && number.ParseFloat(v) >= 500 {
		return color.Brown(v)
	}

	return v
}
