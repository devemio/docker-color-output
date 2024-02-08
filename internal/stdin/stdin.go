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

const rowsCap = 20

// clsBytes contains bytes to clear
// the screen for nix systems.
var clsBytes = []byte("\033[2J\033[H") //nolint:gochecknoglobals

func Get(fn func(rows []string) error) error {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("stdin: %w", err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 && fi.Size() <= 0 {
		return fmt.Errorf("stdin: %w", ErrEmpty)
	}

	rows := make([]string, 0, rowsCap)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row, found := scanner.Bytes(), false //nolint:wastedassign

		if row, found = bytes.CutPrefix(row, clsBytes); found && len(rows) > 0 {
			stdout.Print(string(clsBytes))

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

	if err = fn(rows); err != nil {
		return err
	}

	return nil
}
