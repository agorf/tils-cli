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
go get -u github.com/agorf/tils-cli/cmd/tils
```

**Note:** This command will place the `tils` executable under the _bin_ directory of your `$GOPATH`. To be able to execute it without typing its absolute path, you can update your `PATH` variable to include Go's bin directory by adding the following line to the proper dotfile for your shell (`~/.bashrc`, `~/.zshrc` etc):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Configuration

First [sign up][] for an account and get your [API token][].

[sign up]: https://tils.dev/signup
[API token]: https://tils.dev/account

Use the `config` command to enter your API token:

```plaintext
$ tils config
? API token: 41271f50fd9a7dcf904a727039a5ec24
Wrote config
```

Alternatively, you can set the `TILS_CLI_API_TOKEN` environment variable which
takes precedence over the config file.

## Usage

```plaintext
$ tils help
tils [command]

Commands:

    new        Create til
    show       Show til content in the terminal
    open       Open til in the browser
    copy       Copy til to the clipboard
    edit       Edit til
    archive    Archive til
    delete     Delete til
    config     Configure
    version    Print the current version
    help       Print this help text

If a command is not provided, a picker will ask for one
```

## License

[The MIT License][]

[The MIT License]: https://github.com/agorf/tils-cli/blob/master/LICENSE.txt

## Author

Angelos Orfanakos, <https://angelos.dev>
