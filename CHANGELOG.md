# CHANGELOG

## next

- Addition of a `search` command. This allows you to run JQL either via the configuration file (you specify `--template` to find the correct template in the config file), or the CLI `--query` option. It displays the ID, and the title of the issues it finds. If `fields.story_point_field` is defined in your configuration, the output will also include the Story Point value.

## 2.1.0

- Addition of the `-v` flag on the `sprint list` command.

## 2.0.0

- The configuration file has now moved to `~/.benmatselby/walter.yml`.

## 1.0.0

- First initial versioned release of walter.
- Ability to list boards.
- Ability to list all issues for a given board.
- Ability to list all issues for a given board.
- Ability to list sprints.
- Ability to list issues for a sprint.
- Ability to show a burn down table of data for a given board and sprint.
