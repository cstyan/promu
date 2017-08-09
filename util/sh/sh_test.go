package sh

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestSplitParameters(t *testing.T) {
	in := `-a -tags 'netgo static_build'`
	expect := []string{"-a", "-tags", `'netgo static_build'`}
	got := SplitParameters(in)
	for i, g := range got {
		if expect[i] != g {
			t.Error("expected", expect[i], "got", g, "full output: ", strings.Join(got, "#"))
		}
	}
}

func TestPath(t *testing.T) {
	in := []string{".release", "tarball.tar.gz"}
	expect := fmt.Sprintf(".release%starball.tar.gz", string(filepath.Separator))
	got := Path(in...)
	if got != expect {
		t.Errorf("expected %+v, got %+v", expect, got)
	}
}
