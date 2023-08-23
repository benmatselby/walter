# Walter

![GitHub Badge](https://github.com/benmatselby/walter/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmatselby/walter)](https://goreportcard.com/report/github.com/benmatselby/walter)

_That rug really tied the room together, did it not?_

CLI application for getting information out of Jira.

```text
CLI application for retrieving data from Jira

Usage:
  walter [command]

Available Commands:
  board       Board related commands
  help        Help about any command
  search      Search for issues
  sprint      Sprint related commands

Flags:
      --config string   config file (default is $HOME/.config/benmatselby/walter.yaml)
  -h, --help            help for walter

Use "walter [command] --help" for more information about a command.
```

## Requirements

If you are wanting to build and develop this, you will need the following items installed. If, however, you just want to run the application I recommend using the docker container (See below).

- Go version 1.20+

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

There is also a configuration file (by default, `~/.config/benmatselby/walter.yaml`) that allows you to configure various things:

- The story point field in your Jira instance, per board.
- The layout of the board (Couldn't find an API to get this at the moment).
- Template definitions for the `search` command.

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

fields:
  story_point_field: customfield_10911

templates:
  all-open:
    query: "project = 'My Project' AND status != Completed ORDER BY Summary"
    count: 200
```

| Item                       | Description                                                                                                                                                                                                                                                                                                                                                   |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `boards`                   | Top level node for board configuration.                                                                                                                                                                                                                                                                                                                       |
| `My special board`         | The name of the board (`walter board list`).                                                                                                                                                                                                                                                                                                                  |
| `story_point_field`        | Defines the custom field that is houses the story point estimation (If you do not define this, the sprint.burndown command does not fully render all the information). If you know there is only one field in your system for story points, define this field.                                                                                                |
| `story_point_fields`       | Defines the custom fields that should be used in conjunction with one another for the story point estimation. It turns out that some projects have multiple fields defined over time! _I highly recommend against this_. It will try to use `story_point_field` first, and then iterate every field defined in this config until it gets a story point value. |
| `layout`                   | The names of the columns on the board (I am struggling to find an API endpoint that documents this).                                                                                                                                                                                                                                                          |
| `templates`                | Top level node for template configuration.                                                                                                                                                                                                                                                                                                                    |
| `fields.story_point_field` | Top the overarching story point field to be used. Long term, the commands will look at this value first, before looking at board specific values.                                                                                                                                                                                                             |

## Installation

You can install this application a few ways:

<details>
<summary>Installation via Docker</summary>

Other than requiring [docker](http://docker.com) to be installed, there are no other requirements to run the application this way.

```shell
$ docker build -t benmatselby/walter .
$ docker run \
  --rm \
  -t \
  -eJIRA_TOKEN \
  -eJIRA_URL \
  -eJIRA_USERNAME \
  -v "${HOME}/.config/benmatselby":/root/.config/benmatselby \
  benmatselby/walter:latest "$@"
```

The `latest` tag mentioned above can be changed to a released version. For all releases, see [here](https://hub.docker.com/repository/docker/benmatselby/walter/tags). An example would then be:

```shell
benmatselby/walter:version-2.2.0
```

This would use the `verson-2.2.0` release in the docker command.

**Note** - if you get the following error when running this on Windows using [docker desktop for Windows](https://hub.docker.com/editions/community/docker-ce-desktop-windows):

`Failed to load config: Config File "walter" Not Found in "[/root/.config/benmatselby]"`

Then you may need to reset your credentials in _[docker](http://docker.com) > Settings > Shared Drives > Reset credentials_, because it uses shared drives for mounting volumes and caches Active Directory credentials.

</details>

<details>
<summary>Installation via Git</summary>

```shell
git clone git@github.com:benmatselby/walter.git
cd walter
make all
./walter
```

You can also install into your `$GOPATH/bin` by running `make build && go install`.

</details>

<details>
<summary>Installation via Binaries</summary>

You can download the binaries from the [release pages](https://github.com/benmatselby/walter/releases). Find the release you want, and check the "Assets" section.

Once downloaded, be sure to put the binary in a folder that is referenced in your `$PATH``.

</details>

## Testing

To generate the code used to mock away the Jira interaction, run the following command.

```shell
mockgen -source jira/jira.go
```

This will generate you some source code you can copy into `jira/mock_jira.go`. You will need to change the package to `jira`.
