package stdout

import (
	"fmt"
)

func Print(in string) {
	fmt.Print(in) //nolint:forbidigo
}

func Println(in string) {
	fmt.Println(in) //nolint:forbidigo
}
