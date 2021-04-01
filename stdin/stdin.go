package stdin

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetLines() ([]string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 || fi.Size() <= 0 {
		return nil, errors.New("no stdin")
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
		return nil, err
	}

	return lines, nil
}
