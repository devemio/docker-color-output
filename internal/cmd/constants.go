package cmd

import "docker-color-output/internal/lines/fmt"

type Cmd string

func (c Cmd) String() string {
	return string(c)
}

func (c Cmd) Columns() []string {
	switch c {
	case DockerPs:
		return []string{
			"CONTAINER ID", //
			"IMAGE",        //
			"COMMAND",      //
			"CREATED AT",   // opt
			"CREATED",      //
			"PORTS",        // nullable
			"STATE",        // opt
			"STATUS",       //
			"SIZE",         // opt
			"NAMES",        //
			"LABELS",       // opt
			"MOUNTS",       // opt | nullable
			"NETWORKS",     // opt | nullable
		}
	case DockerImages:
		return []string{
			"IMAGE ID",   //
			"REPOSITORY", //
			"TAG",        //
			"DIGEST",     // opt
			"CREATED",    //
			"CREATED AT", // opt
			"SIZE",       //
		}
	case DockerComposePs:
		return []string{
			"NAME",
			"COMMAND",
			"SERVICE",
			"STATUS",
			"PORTS", // nullable
		}
	default:
		return nil
	}
}

func (c Cmd) GetFmt() fmt.Formatable {
	switch c {
	case DockerPs:
		return fmt.NewDockerPsLineFmt()
	case DockerImages:
		return fmt.NewDockerImagesLineFmt()
	case DockerComposePs:
		return fmt.NewDockerComposePsLineFmt()
	default:
		return nil
	}
}

const (
	DockerPs        Cmd = "docker ps"
	DockerImages    Cmd = "docker images"
	DockerComposePs Cmd = "docker compose ps"
)
