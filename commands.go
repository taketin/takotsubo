package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/taketin/takotsubo/gh"
)

var Commands = []cli.Command{
	commandGetteam,
	commandAdd,
	commandMembers,
}

var commandGetteam = cli.Command{
	Name:  "getteam",
	Usage: "getteam {repository}",
	Description: `
	Get teams information from repository.
`,
	Action: doGetteam,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "add {team_id} {member,member,member}",
	Description: `
	Adding member to team.
`,
	Action: doAdd,
}

var commandMembers = cli.Command{
	Name:  "members",
	Usage: "members {team_id}",
	Description: `
	Show team members.
`,
	Action: doMembers,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doGetteam(c *cli.Context) {
	repo := c.Args()[0]
	teams := gh.Teams(repo)

	for _, team := range teams {
		fmt.Printf("ID: %d NAME: %s URL: %s\n", team.Id, team.Name, team.Url)
	}
}

func doAdd(c *cli.Context) {
	teamId := c.Args()[0]
	membersCsv := c.Args()[1]

	members := strings.Split(membersCsv, ",")
	if gh.AddMembers(teamId, members) {
		fmt.Print("Succeeded !")
	} else {
		fmt.Print("Failed! Sorry... ><")
	}
}

func doMembers(c *cli.Context) {
	teamId := c.Args()[0]
	members := gh.Members(teamId)

	for _, member := range members {
		fmt.Printf("%s\n", member)
	}
}
