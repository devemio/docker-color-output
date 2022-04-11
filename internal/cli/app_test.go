package cli

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"docker-color-output/internal/lines"
	"docker-color-output/pkg/utils/assert"
)

func TestExecute(t *testing.T) {
	read := func(filename string) []string {
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(b)
		bytes, _ := os.ReadFile(path + "/../../test/data/" + filename)
		vals := strings.Split(string(bytes), "\n")
		vals = vals[:len(vals)-1]
		if len(vals) == 0 {
			return nil
		}
		return vals
	}

	build := func(in []*lines.Line) []string {
		var vals []string
		for _, v := range in {
			vals = append(vals, v.Build())
		}
		return vals
	}

	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"no_stdin":                {wantErr: true},
		"no_first_line":           {in: "in/no_first_line.out", wantErr: true},
		"invalid_cols":            {in: "in/invalid_cols.out", wantErr: true},
		"docker_ps":               {in: "in/docker_ps.out", want: "out/docker_ps.out"},
		"docker_ps:custom_cols":   {in: "in/docker_ps_custom_cols.out", want: "out/docker_ps_custom_cols.out"},
		"docker_ps:nullable_col":  {in: "in/docker_ps_nullable_col.out", want: "out/docker_ps_nullable_col.out"},
		"docker_ps:nullable_cols": {in: "in/docker_ps_nullable_cols.out", wantErr: true},
		"docker_images":           {in: "in/docker_images.out", want: "out/docker_images.out"},
		"docker_compose_ps":       {in: "in/docker_compose_ps.out", want: "out/docker_compose_ps.out"},
		"docker_compose_ps_v1":    {in: "in/docker_compose_ps_v1.out", want: "out/docker_compose_ps_v1.out"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			vals, err := Execute(read(tt.in))
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, read(tt.want), build(vals))
		})
	}
}
