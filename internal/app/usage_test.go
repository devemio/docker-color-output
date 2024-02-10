package app_test

import (
	"errors"
	"testing"

	"github.com/devemio/docker-color-output/internal/app"
)

var errTest = errors.New("test")

func TestUsage(*testing.T) {
	app.Usage(errTest)
}
