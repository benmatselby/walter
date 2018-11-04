# Walter

[![Build Status](https://travis-ci.org/benmatselby/walter.png?branch=master)](https://travis-ci.org/benmatselby/walter)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmatselby/walter?style=flat-square)](https://goreportcard.com/report/github.com/benmatselby/walter)

_That rug really tied the room together, did it not?_

CLI application for getting information out of Jira. It's based on [Donny](https://github.com/benmatselby/donny) so the aims are the same:

```
CLI application for retrieving data from Jira

Usage:
  walter [command]

Available Commands:
  boards      Board related commands
  help        Help about any command
  sprint      Sprint related commands

Flags:
      --config string   config file (default is $HOME/.walter/config.yaml)
  -h, --help            help for walter
  -t, --toggle          Help message for toggle

Use "walter [command] --help" for more information about a command.
```

## Requirements

If you are wanting to build and develop this, you will need the following items installed. If, however, you just want to run the application I recommend using the docker container (See below)

* Go version 1.11+
* [Dep installed](https://github.com/golang/dep)

## Configuration

You will need the following environment variables defining:

```bash
export JIRA_TOKEN=""
export JIRA_URL=""
export JIRA_USERNAME=""
```

## Installation via Docker

Other than requiring [docker](http://docker.com) to be installed, there are no other requirements to run the application this way. This is the preferred method of running the `walter`. The image is [here](https://hub.docker.com/r/benmatselby/walter/).

```bash
$ docker run \
    --rm \
    -t \
    -eJIRA_TOKEN \
    -eJIRA_URL \
    -eJIRA_USERNAME \
    benmatselby/walter "$@"
```

## Installation via Git

```bash
git clone git@github.com:benmatselby/walter.git
cd walter
make all
./walter
```
