package client

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/docopt/docopt-go"
)

func (c *DeisClient) AuthRegister(args []string) (err error) {
	Usage := `
Register a new user with a Deis controller

Usage: deis auth:register <controller> [options]

Options:
  --username=USERNAME    provide a username for the new account
  --password=PASSWORD    provide a password for the new account
  --email=EMAIL          provide an email address
`
	parsed, err := docopt.Parse(Usage, args, true, "", false, false)
	if err != nil {
		return err
	}
	var controller, username, password, email string
	controllerURL, err := url.Parse(parsed["<controller>"].(string))
	if err != nil {
		return err
	}
	if controllerURL.Scheme == "" {
		controllerURL.Scheme = "http"
	}
	controller = controllerURL.String()
	if parsed["--username"] != nil {
		username = parsed["--username"].(string)
	} else {
		fmt.Print("username: ")
		fmt.Scanf("%s", &username)
	}
	if parsed["--password"] != nil {
		password = parsed["--password"].(string)
	} else {
		fmt.Print("password: ")
		fmt.Scanf("%s", &password)
		var passwordConfirm string
		fmt.Print("password (confirm): ")
		fmt.Scanf("%s", &passwordConfirm)
		if password != passwordConfirm {
			return errors.New("password mismatch, aborting registration")
		}
	}
	if parsed["--email"] != nil {
		email = parsed["--email"].(string)
	}
	fmt.Println(controller, username, password, email)
	// // fmt.Println(parsed)
	// // if not urlparse.urlparse(controller).scheme:
	// // 	controller = "http://{}".format(controller)
	//
	// fmt.Println(url)
	return nil
}
