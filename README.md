# tils-cli

![](https://img.shields.io/github/v/tag/agorf/tils-cli?label=version&sort=semver)

A command-line client for [tils.dev][]

[tils.dev]: https://tils.dev/

**Note:** Visit the [master branch][master] for the latest released version.

[master]: https://github.com/agorf/tils-cli/tree/master

## Installation

### Pre-built binaries

Binaries for Linux, MacOS and Windows are available under [releases][].

[releases]: https://github.com/agorf/tils-cli/releases

### From source

You need to have [Go](https://golang.org/) installed.

On Debian GNU/Linux:

```shell
sudo apt install golang
```

On MacOS:

```shell
brew install go
```

Then issue:

```shell
go get github.com/agorf/tils-cli/cmd/tils
```

## Configuration

First [sign up][] for an account and get your [API token][].

[sign up]: https://tils.dev/signup
[API token]: https://tils.dev/account

Create a file under `~/.config/tils-cli.json` and enter:

```json
{
  "api_token": "mytoken"
}
```

Replace `mytoken` with your API token.

Alternatively, you can set the `TILS_CLI_API_TOKEN` environment variable which
takes precedence over the config file.

## Usage

This is the output of the `help` command:

```plaintext
tils [command]

Commands:

    new        Create til
    show       Show til content in the terminal
    open       Open til in the browser
    copy       Copy til to the clipboard
    edit       Edit til
    archive    Archive til
    delete     Delete til
    version    Print the current version
    help       Print this help text

If a command is not provided, a picker will ask for one
```

## License

[The MIT License][]

[The MIT License]: https://github.com/agorf/tils-cli/blob/master/LICENSE.txt

## Author

Angelos Orfanakos, <https://angelos.dev>
