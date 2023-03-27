package stdin

import (
	"bufio"
	"errors"
	"os"
)

var ErrNoStdin = errors.New("no stdin")

func Get() ([]string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	if fi.Mode()&os.ModeNamedPipe == 0 && fi.Size() <= 0 {
		return nil, ErrNoStdin
	}

	var res []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		res = append(res, s.Text())
	}

	if err = s.Err(); err != nil {
		return nil, err //nolint:wrapcheck
	}

	return res, nil
}
