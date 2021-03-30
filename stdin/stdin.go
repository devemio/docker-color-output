package stdin

import (
	"bufio"
	"github.com/devemio/docker-color-output/console"
	"os"
	"strings"
)

func GetLines() []string {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		console.Usage()
		os.Exit(1)
	}

	var lines []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "-") {
			lines = append(lines, s.Text())
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return lines
}
