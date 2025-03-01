package stderr

import (
	"fmt"
	"os"
)

func Println(in string) {
	_, _ = fmt.Fprintln(os.Stderr, in)
}
