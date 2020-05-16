# tils-cli

![](https://img.shields.io/github/v/tag/agorf/tils-cli?label=version&sort=semver)

**tils-cli** is a command-line client for [tils.dev][]

[tils.dev]: https://tils.dev/

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
go get https://github.com/agorf/tils-cli
```

## Configuration

You need to set the `TILS_CLI_API_TOKEN` environment variable to your [tils.dev][] API token.

To get your [tils.dev][] API token, [sign up][] and go to [your account][account].

[sign up]: https://tils.dev/signup
[account]: https://tils.dev/account

## Usage

This is the output of `tils-cli help`:

```plaintext
tils-cli [command]

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
