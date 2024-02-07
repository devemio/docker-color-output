package app_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/devemio/docker-color-output/internal/app"
	"github.com/devemio/docker-color-output/pkg/util/assert"
)

func TestRun(t *testing.T) {
	t.Parallel()

	read := func(filename string) []string {
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(b)
		bytes, _ := os.ReadFile(path + "/../../test/data/" + filename)
		res := strings.Split(string(bytes), "\n")
		res = res[:len(res)-1]

		if len(res) == 0 {
			return nil
		}

		return res
	}

	//nolint:exhaustruct
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"no_stdin":                {wantErr: true},
		"no_first_line":           {in: "no_first_line.out", wantErr: true},
		"invalid_cols":            {in: "invalid_cols.out", wantErr: true},
		"docker_ps":               {in: "docker_ps.out", want: "docker_ps.out"},
		"docker_ps:custom_cols":   {in: "docker_ps_custom_cols.out", want: "docker_ps_custom_cols.out"},
		"docker_ps:nullable_col":  {in: "docker_ps_nullable_col.out", want: "docker_ps_nullable_col.out"},
		"docker_ps:nullable_cols": {in: "docker_ps_nullable_cols.out", wantErr: true},
		"docker_images":           {in: "docker_images.out", want: "docker_images.out"},
		"docker_compose_ps":       {in: "docker_compose_ps.out", want: "docker_compose_ps.out"},
		"docker_compose_ps_1":     {in: "docker_compose_ps_1.out", want: "docker_compose_ps_1.out"},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rows, err := app.Run(read("in/" + tt.in))
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, read("out/"+tt.want), rows)
		})
	}
}
