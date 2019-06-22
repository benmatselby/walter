# Walter

[![Build Status](https://travis-ci.org/benmatselby/walter.png?branch=master)](https://travis-ci.org/benmatselby/walter)
[![codecov](https://codecov.io/gh/benmatselby/walter/branch/master/graph/badge.svg)](https://codecov.io/gh/benmatselby/walter)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmatselby/walter)](https://goreportcard.com/report/github.com/benmatselby/walter)

_That rug really tied the room together, did it not?_

CLI application for getting information out of Jira. It's based on [Donny](https://github.com/benmatselby/donny) so the aims are the same:

```text
CLI application for retrieving data from Jira

Usage:
  walter [command]

Available Commands:
  board       Board related commands
  help        Help about any command
  sprint      Sprint related commands

Flags:
      --config string   config file (default is $HOME/.walter/config.yaml)
  -h, --help            help for walter
  -t, --toggle          Help message for toggle

Use "walter [command] --help" for more information about a command.
```

## Requirements

If you are wanting to build and develop this, you will need the following items installed. If, however, you just want to run the application I recommend using the docker container (See below).

- Go version 1.11+

## Configuration

### Environment variables

You will need the following environment variables defining:

```shell
export JIRA_TOKEN=""
export JIRA_URL=""
export JIRA_USERNAME=""
```

Creating a Jira API Token is documented [here](https://confluence.atlassian.com/cloud/api-tokens-938839638.html).

### Application configuration file

Long term this may not be required, but right now we need a configuration file (by default, `~/.walter/config.yaml`).

An example:

```yml
boards:
  My special board:
    story_point_field: customfield_888
    story_point_fields:
      - customfield_888
      - customfield_889
    layout:
      - To Do
      - In Progress
      - Stuck
      - Review
      - Done
```

| Item                 | Description                                                                                                                                                                                                                                                                                                                                                        |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `boards`             | This is the top level node for board configuration.                                                                                                                                                                                                                                                                                                                |
| `My special board`   | This is the name of the board (`walter board list`).                                                                                                                                                                                                                                                                                                               |
| `story_point_field`  | This defines the custom field that is houses the story point estimation (If you do not define this, the sprint.burndown command does not fully render all the information). If you know there is only one field in your system for story points, define this field.                                                                                                |
| `story_point_fields` | This defines the custom fields that should be used in conjunction with one another for the story point estimation. It turns out that some projects have multiple fields defined over time! _I highly recommend against this_. It will try to use `story_point_field` first, and then iterate every field defined in this config until it gets a story point value. |
| `layout`             | This is the names of the columns on the board (I am struggling to find an API endpoint that documents this).                                                                                                                                                                                                                                                       |

## Installation via Docker

Other than requiring [docker](http://docker.com) to be installed, there are no other requirements to run the application this way. This is the preferred method of running the `walter`. The image is [here](https://hub.docker.com/r/benmatselby/walter/).

```shell
$ docker run \
    --rm \
    -t \
    -eJIRA_TOKEN \
    -eJIRA_URL \
    -eJIRA_USERNAME \
    -v "${HOME}/.walter":/root/.walter \
    benmatselby/walter "$@"
```

## Installation via Git

```shell
git clone git@github.com:benmatselby/walter.git
cd walter
make all
./walter
```

You can also install into your `$GOPATH/bin` by running `make build && go install`.

## Testing

To generate the code used to mock away the Jira interaction, run the following command.

```shell
mockgen -source jira/jira.go
```

This will generate you some source code you can copy into `jira/mock_jira.go`. You will need to change the package to `jira`.
