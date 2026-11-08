// Command gyro is the official Gyro Hosting command-line interface.
package main

import (
	"os"

	"github.com/Gyro-Hosting/gyro-cli/internal/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.Version, cli.Commit, cli.Date = version, commit, date
	os.Exit(cli.Run(os.Args[1:]))
}
