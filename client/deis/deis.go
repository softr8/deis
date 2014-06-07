package main

import (
	"github.com/deis/deis/client"
	"os"
)

func main() {
	os.Exit(client.Command(nil))
}
