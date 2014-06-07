package client

import (
	"fmt"
	"strings"

	"github.com/docopt/docopt-go"
)

// Shortcuts are terse versions of longer commands.
var SHORTCUTS = [][]string{
	{"create", "apps:create"},
	{"destroy", "apps:destroy"},
	{"init", "clusters:create"},
	{"info", "apps:info"},
	{"run", "apps:run"},
	{"open", "apps:open"},
	{"logs", "apps:logs"},
	{"register", "auth:register"},
	{"login", "auth:login"},
	{"logout", "auth:logout"},
	{"ps", "containers:list"},
	{"scale", "containers:scale"},
	{"rollback", "releases:rollback"},
	{"sharing", "perms:list"},
	{"sharing:list", "perms:list"},
	{"sharing:add", "perms:create"},
	{"sharing:remove", "perms:delete"},
}

func (c *DeisClient) Shortcuts(args []string) (err error) {
	Usage := `
Show valid shortcuts for client commands.

Usage: deis shortcuts
`
	_, _ = docopt.Parse(Usage, args, true, "", false)
	fmt.Print("Valid shortcuts are:\n\n")
	for _, shortcut := range SHORTCUTS {
		short, cmd := shortcut[0], shortcut[1]
		if !strings.Contains(short, ":") {
			fmt.Printf("%-10v -> %v\n", short, cmd)
		}
	}
	fmt.Println("\nUse \"deis help [command]\" to learn more")
	return
}
