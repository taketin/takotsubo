package gh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var ApiVersion string // ex. "/api/v3"
var Token string
var Host string

var exit = os.Exit

type Member struct {
	Login string
}

type Team struct {
	Url string
	Name string
	Id int
}

func AddMembers(teamId string, members []string) bool {
	prepare()

	succeeded := false
	url := Host + ApiVersion + "/teams/" + teamId + "/members/"

	for _, member := range members {
		req, _ := http.NewRequest("PUT", url + member, nil)
		body := performRequest(req)
		if len(body) == 0 {
			succeeded = true
		}
	}

	return succeeded
}

func Teams(repo string) []Team {
	prepare()

	var teams []Team

	url := Host + ApiVersion + "/orgs/" + repo + "/teams"
	req, _ := http.NewRequest("GET", url, nil)
	body := performRequest(req)

	teamsJson := make([]Team, 0)
	_ = json.Unmarshal(body, &teamsJson)

	for _, team := range teamsJson {
		teams = append(teams, team)
	}

	return teams
}

func Members(teamId string) []string {
	prepare()

	var members []string

	url := Host + ApiVersion + "/teams/" + teamId + "/members"
	req, _ := http.NewRequest("GET", url, nil)
	body := performRequest(req)

	membersJson := make([]Member, 0)
	_ = json.Unmarshal(body, &membersJson)

	for _, member := range membersJson {
		members = append(members, member.Login)
	}

	return members
}

func prepare() bool {
	succeeded := true
	if Token == ""  || Host == "" || ApiVersion == "" {
		println("Please settings for 'Token', 'Host' and 'ApiVersion'")
		succeeded = false
		exit(1)
	}

	return succeeded
}

func performRequest(req *http.Request) []uint8 {
	req.SetBasicAuth(Token, "x-oauth-basic")
	client := http.Client{}
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	return body
}
