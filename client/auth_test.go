package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthRegister(t *testing.T) {
	// test the usage string
	args := []string{"auth:register"}
	out := commandOutput(args)
	if !strings.HasPrefix(out, "Usage: deis auth:register <controller>") {
		t.Error(out)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake twitter json string"}`)
	}))
	defer ts.Close()
	// test registration
	args = []string{
		"register", ts.URL,
		"--username=test", "--password=asdf1234", "--email=test@deis.io",
	}
	out = commandOutput(args)
	fmt.Println(out)
	if !strings.HasPrefix(out, "Use: deis auth:register <controller>") {
		t.Error(out)
	}
}
