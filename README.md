takotsubo
====

## Description

GitHub Team Inviter

## Usage

```bash
$ takotsubo getteam {repository name}
> ID: 001 NAME: Owners URL: http://github.com/xxxxx
> ID: 002 NAME: Engineers URL: http://github.com/xxxxx

$ takotsubo add {team id} {member_name}

// If you want to specify multiple members
$ takotsubo add {team id} {member_name,member_name,member_name...}

$ takotsubo members {team id}
```

## Install

To install, use `go get`:

```bash
$ go get github.com/taketin/takotsubo

// Settings for Environment variables
$ export GITHUB_TOKEN="your token (required ADMIN authority)"

// If your want to use in the GH:E
$ export GITHUB_HOST="your GitHub Enterprise host"
```

## Contribution

1. Fork ([https://github.com//takotsubo/fork](https://github.com//takotsubo/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[@taketin](https://github.com/taketin)
