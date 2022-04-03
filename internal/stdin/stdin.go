package stdin

import (
	"bufio"
	"errors"
	"os"
)

func Get() ([]string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 && fi.Size() <= 0 {
		return nil, errors.New("no stdin")
	}

	var res []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		res = append(res, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return res, nil
}