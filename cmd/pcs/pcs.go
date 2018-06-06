package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shuvigoss/pcs/config"
	"github.com/shuvigoss/pcs/exec"
)

var configPath = flag.String("config", "pcs.json", "config file path")
var command = flag.String("command", "whoami", "command to exec")

func main() {
	flag.Usage = func() { usageExit(0) }
	flag.Parse()

	c := config.NewConfig()
	c.ParseFile(*configPath)

	exec.RunCommands(*c, *command)
}
func usageExit(i int) {
	fmt.Println("./pcs -config=./pcs.json -command=\"whoami\"")
	os.Exit(i)
}
