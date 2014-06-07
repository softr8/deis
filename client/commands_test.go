package client

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func commandOutput(args []string) (output string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Command(args)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	output = <-outC
	return
}

func TestHelp(t *testing.T) {
	allArgs := [][]string{{"-h"}, {"--help"}, {"help"}}
	out := ""
	for _, args := range allArgs {
		out = commandOutput(args)
		if !strings.HasPrefix(out, "The Deis command-line client") ||
			!strings.HasSuffix(out, "to deploy to an application.\n") {
			t.Error(out)
		}
	}
}
func TestUsage(t *testing.T) {
	out := commandOutput(nil)
	if out != "Usage: deis <command> [<args>...]\n" {
		t.Error(out)
	}
}

func TestVersion(t *testing.T) {
	args := []string{"--version"}
	out := commandOutput(args)
	if out != "Deis CLI 0.10.0\n" {
		t.Error(out)
	}
}
