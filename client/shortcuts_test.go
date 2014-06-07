package client

import (
	"strings"
	"testing"
)

func TestShortcuts(t *testing.T) {
	args := []string{"shortcuts"}
	out := commandOutput(args)
	if !strings.HasPrefix(out, "Valid shortcuts are:") {
		t.Error(out)
	}
}
