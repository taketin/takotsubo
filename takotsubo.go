package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/taketin/takotsubo/gh"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	host := os.Getenv("GITHUB_HOST")

	if token == "" {
		println("Please setting environment variables for 'GITHUB_TOKEN' and 'GITHUB_HOST'")
		os.Exit(1)
	}

	if host == "" {
		host = "https://github.com"
	}

	gh.Token = token
	gh.Host = host
	gh.ApiVersion = "/api/v3"

	app := cli.NewApp()
	app.Name = "takotsubo"
	app.Version = Version
	app.Usage = ""
	app.Author = "taketin"
	app.Email = "tksthdnr@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}
