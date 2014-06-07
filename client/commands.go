package client

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

type DeisClient struct {
}

// Command executes a given command line
func Command(argv []string) (returnCode int) {
	Usage := `
The Deis command-line client issues API calls to a Deis controller.

Usage: deis <command> [<args>...]

Auth commands:

  register      register a new user with a controller
  login         login to a controller
  logout        logout from the current controller

Subcommands, use "deis help [subcommand]" to learn more:

  formations    manage formations used to host applications
  layers        manage layers used for node configuration
  nodes         manage nodes used to host containers and proxies

  apps          manage applications used to provide services
  containers    manage containers used to handle requests and jobs
  config        manage environment variables that define app config
  builds        manage builds created using "git push"
  releases      manage releases of an application

  providers     manage credentials used to access cloud providers
  flavors       manage flavors of nodes including size and location
  keys          manage ssh keys used for "git push" deployments

  perms         manage permissions for shared apps and formations

Developer shortcut commands:

  create        create a new application
  scale         scale containers by type (web=2, worker=1)
  info          view information about the current app
  open          open a URL to the app in a browser
  logs          view aggregated log info for the app
  run           run a command in an ephemeral app container
  destroy       destroy an application

Use "git push deis master" to deploy to an application.

`
	returnCode = 0
	argv = parseArgs(argv)
	args, err := docopt.Parse(
		Usage, argv, true, "Deis CLI 0.10.0", true, false)
	if err != nil {
		returnCode = 1
		return
	} else if len(args) == 0 {
		return
	}
	cmd := args["<command>"].(string)

	cli := DeisClient{}

	switch cmd {
	// case "auth:login":
	// 	err = cli.AuthLogin(argv)
	// case "auth:logout":
	// 	err = cli.AuthLogout(argv)
	case "auth:register":
		err = cli.AuthRegister(argv)
	// case "keys:add":
	// 	err = cli.KeysAdd(argv)
	case "shortcuts":
		err = cli.Shortcuts(argv)
	default:
		fmt.Println("Found no matching command, try `deis help`")
		returnCode = 1
	}

	if err != nil {
		returnCode = 1
		fmt.Println(err)
	}

	return
}

func parseArgs(argv []string) []string {
	if argv == nil {
		argv = os.Args[1:]
	}

	if len(argv) > 0 {
		// parse "deis help <command>" as "deis <command> --help"
		if argv[0] == "help" {
			argv = append(argv[1:], "--help")
		} else {
			// replace short version of command with full one, if needed
			for _, shortcut := range SHORTCUTS {
				short, full := shortcut[0], shortcut[1]
				if argv[0] == short {
					argv[0] = full
					break
				}
			}
		}
	}
	return argv
}
