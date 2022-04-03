package stdout

import (
	"fmt"

	"docker-color-output/internal/lines"
)

func Println(v lines.Buildable) {
	fmt.Println(v.Build())
}
