# CHANGELOG

## next

- Bumped Go version to 1.23 runtime.

## 2.8.0

- Allow the caller of `make docker-build` to specify the Docker platform.
- Add in the release version to the auto-generated binaries during a release.

## 2.7.0

- Bumped docker image to Go 1.21 runtime.
- Display story point values on the `walter board issues [board]` command.

## 2.6.0

- Bumped docker image to Go 1.19 runtime.

## 2.5.0

- Provide the ability to `--group-by` on the `sprint issues` command.
  - You can group by `board` or `label`. This will provide a list of issues either under the board column heading, or under each label an issue has assigned to it.
  - Note that if you group by `label` an issue may appear in the output twice, under each label.
- Consistently render the issue (including story points if configured) in the search results.
- The releases will now include the built binaries for different platforms/architectures.

## 2.4.0

- Bumped docker image to Go 1.18 runtime.
- Show story point values in the `sprint issues` output, if the field is defined.

## 2.3.0

- Bumped docker image to Go 1.16 runtime.
- Bump the build environment to test on 1.16.

## 2.2.1

- Bump the dependencies.
- Run the GitHub actions on multiple versions of Go.
- List all the boards, not just the limit of 50. Thanks to [Richard Neal](https://github.com/Richard-W-Neal) for raising.
  - `walter board list`
- Use the `Name` attribute to search for a single board, rather than pulling all boards back, and iterating over them, to find the match.

## 2.2.0

- Addition of a `search` command. This allows you to run JQL either via the configuration file (you specify `--template` to find the correct template in the config file), or the CLI `--query` option. It displays the ID, and the title of the issues it finds. If `fields.story_point_field` is defined in your configuration, the output will also include the Story Point value.
- Option of `--format` on the search command, that allows you to either get a list of issues, or a table view of data. When coupled with the story point field, you can see a tablular view of: Total issues, Total story points, Total not story pointed.

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
