package stdin

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/devemio/docker-color-output/internal/stdout"
)

var ErrEmpty = errors.New("empty")

const (
	capRows = 50
)

//nolint:gochecknoglobals
var (
	xClearToEnd     = []byte("\033[J")
	xClearToLineEnd = []byte("\033[K")
	xMoveCursorHome = []byte("\033[H")
)

//nolint:cyclop
func Get(fn func(rows []string) error) error {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("stdin: %w", err)
	} else if fi.Mode()&os.ModeNamedPipe == 0 && fi.Size() <= 0 {
		return fmt.Errorf("stdin: %w", ErrEmpty)
	}

	var (
		rows    = make([]string, 0, capRows)
		scanner = bufio.NewScanner(os.Stdin)
	)

	for scanner.Scan() {
		//nolint:wastedassign
		row, found := scanner.Bytes(), false

		// xClearToLineEnd
		if _, found = bytes.CutPrefix(row, xClearToLineEnd); found {
			stdout.Print(string(xClearToLineEnd))

			continue
		}

		// xClearToLineEnd
		row = bytes.TrimSuffix(row, xClearToLineEnd)

		// xClearToEnd
		if row, found = bytes.CutPrefix(row, xClearToEnd); found {
			stdout.Print(string(xClearToEnd))
		}

		// xMoveCursorHome
		if row, found = bytes.CutPrefix(row, xMoveCursorHome); found && len(rows) > 0 {
			stdout.Print(string(xMoveCursorHome))

			if err = fn(rows); err != nil {
				return err
			}

			rows = rows[:0]
		}

		rows = append(rows, string(row))
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("stdin: scan: %w", err)
	}

	return fn(rows)
}
