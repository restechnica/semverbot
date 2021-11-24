# Semverbot

[![github.com release badge](https://img.shields.io/github/release/restechnica/semverbot.svg)](https://github.com/restechnica/semverbot/)
[![github.com workflow badge](https://github.com/restechnica/semverbot/workflows/main/badge.svg)](https://github.com/restechnica/semverbot/actions?query=workflow%3Amain)
[![go.pkg.dev badge](https://pkg.go.dev/badge/github.com/restechnica/semverbot)](https://pkg.go.dev/github.com/restechnica/semverbot)
[![goreportcard.com badge](https://goreportcard.com/badge/github.com/restechnica/semverbot)](https://goreportcard.com/report/github.com/restechnica/semverbot)
[![img.shields.io MPL2 license badge](https://img.shields.io/github/license/restechnica/semverbot)](./LICENSE)

A CLI which automates semver versioning.

## Table of Contents

* [Requirements](#requirements)
* [How to install](#how-to-install)
  * [Github](#github)
  * [Homebrew](#homebrew)
* [Usage](#usage)
* [Modes](#modes)
* [How to configure](#how-to-configure)
* [Configuration properties](#configuration-properties)
* [Examples](#examples)
* [Why Semverbot?](#why-semverbot)
  
## Requirements

`sbot` requires a `git` installation.

## How to install

`sbot` can be retrieved from GitHub or a Homebrew tap. Run `sbot -h` to validate the installation.
The tool is available for Windows, Linux and macOS.

### github

`sbot` is available through github. The following example works for a GitHub Workflow, other CI/CD tooling will require a different path setup.

```shell
SEMVERBOT_VERSION=1.0.0
mkdir bin
echo "$(pwd)/bin" >> $GITHUB_PATH
curl -o bin/sbot -L https://github.com/restechnica/semverbot/releases/download/v$SEMVERBOT_VERSION/sbot-linux-amd64
chmod +x bin/sbot
```

### homebrew

`sbot` is available through the public tap [github.com/restechnica/homebrew-tap](https://github.com/restechnica/homebrew-tap)

```shell
brew tap restechnica/tap git@github.com:restechnica/homebrew-tap.git
brew install restechnica/tap/semverbot
```

### golang

`sbot` is written in golang, which means you can use `go install`. Make sure the installation folder, which depends on your golang setup, is in your system PATH.

```shell
go install github.com/restechnica/semverbot/cmd/sbot@v1.0.0
```

## Usage

Each command has a `-h, --help` flag available.

### `sbot get version`

Gets the current version, which is the latest `git` annotated tag without any prefix.

### `sbot init`

Generates a configuration with defaults, see [configuration defaults](#defaults).

### `sbot predict version [-m, --mode] <mode>`

Gets the next version, without any prefix. Uses a mode to detect which semver level it should increment. Defaults to mode `auto`.
See [Modes](#modes) for more documentation on the supported modes.

### `sbot push version`

Pushes the latest `git` tag to the remote repository. Equivalent to `git push origin {prefix}{version}`.

### `sbot release version [-m, --mode] <mode>`

Creates a new version, which is a `git` annotated tag. Uses a mode to detect which semver level it should increment.
Defaults to mode `auto`. See [Modes](#modes) for more documentation on the supported modes.

### `sbot update version`

Fetches all tags with `git` to make sure the git repo has the latest tags available.
Equivalent to running `git fetch --unshallow` and `git fetch --tags`.
This command is very useful in pipelines where shallow clones are often the default to save time and space.

## Modes

### auto (default)

Attempts a series of modes in the following order:
1. `git-branch`
1. `git-commit` - only if `git-branch` failed to detect a semver level to increment
1. `patch` - only if `git-commit` failed to detect a semver level to increment

### git-branch

Detects which semver level to increment based on the **name** of the `git` branch from where a merge commit originated from.
This only works when the old branch has not been deleted yet.

The branch name is matched against the ['semver' configuration](#defaults).

### git-commit

Detects which semver level to increment based on the **message** of the latest `git` commit.

The commit message is matched against the ['semver' configuration](#defaults).

### major

Increments the `major` level.

### minor

Increments the `minor` level.

### patch

Increments the `patch` level.

## How to configure

`sbot` supports a configuration file. It looks for a `.semverbot.toml` file in the current working directory by default.
`.json` and `.yaml` formats are not officially supported, but might work. Using `.toml` is highly recommended.

### Defaults

`sbot init` generates the following configuration:

```toml
mode = "auto"

[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "v"

[semver]
patch = ["fix", "bug"]
minor = ["feature"]
major = ["release"]

[modes]

[modes.git-branch]
delimiters = "/"

[modes.git-commit]
delimiters = "[]"
```

## Configuration properties

### mode

`sbot` supports multiple modes to detect which semver level it should increment. Each mode works with different criteria.
A `mode` flag enables you to switch modes on the fly.

See [Modes](#modes) for documentation about the supported modes.

Defaults to `auto`.

### git

`sbot` works with `git` under the hood, which needs to be set up properly. These config options make sure `git` is set up properly for your environment before running an `sbot` command. 

### git.email

`git` requires `user.email` to be set. If not set, `sbot` will set `user.email` to the value of this property. Rest assured, `sbot` will not override an existing `user.email` value.

Without this config `sbot` might show unexpected behaviour.

### git.name

`git` requires `user.name` to be set. If not set, `sbot` will set `user.name` to the value of this property. Rest assured, `sbot` will not override an existing `user.name` value.

Without this config `sbot` might show unexpected behaviour.

### git.tags.prefix

Different platforms and environments work with different (or without) version prefixes. This option enables you to set whatever prefix you would like to work with.
The `"v"` prefix, e.g. `v1.0.1` is used by default due to its popularity, e.g. some Golang tools completely depend on it.

Note: `sbot` will always display the version without the prefix.

### semver

This is where you configure what you think a semver level should be mapped to.

A mapping of semver levels and words, which are matched against git information.
Whenever a match happens, `sbot` will increment the corresponding level.

See [Modes](#modes) for documentation about the supported modes.

### modes

`sbot` works with different modes, which might require configuration.

### modes.git-branch.delimiters

A string of delimiters which are used to split a git branch name.
The matching words for each semver level in the semver map are matched against each of the resulting strings from the split.

e.g. delimiters `"/"` will split `feature/some-feature` into `["feature", "some-feature"]`,
and the `feature` and `some-feature` strings will be matched against semver map values.

Defaults to `"/"` due to its popular use in git branch names.

### modes.git-commit.delimiters

A string of delimiters which are used to split a git commit message.

e.g. delimiters `"[]"` will split `[feature] some-feature` into `["feature", " some-feature"]`,
and the `feature` and ` some-feature` strings will be matched against semver map values.

Defaults to `"[]"` due to its popular use in git commit messages.

## Examples

### Local

Make sure `sbot` is installed.

```shell
sbot init
sbot release version
sbot push version
```

These commands are basically all you need to work with `sbot` locally.

### GitHub Workflow

#### Shell

```shell
# installation
SEMVERBOT_VERSION=1.0.0
mkdir bin
echo "$(pwd)/bin" >> $GITHUB_PATH
curl -o bin/sbot -L https://github.com/restechnica/semverbot/releases/download/v$SEMVERBOT_VERSION/sbot-linux-amd64
chmod +x bin/sbot

# preparation
sbot update version
echo "CURRENT_VERSION=$(sbot get version)" >> $GITHUB_ENV
echo "RELEASE_VERSION=$(sbot predict version)" >> $GITHUB_ENV
echo "current version: $CURRENT_VERSION"
echo "next version: $RELEASE_VERSION"

# usage
sbot release version
sbot push version
```

#### Yaml

```yaml
name: main

on:
  push:
    branches: [ main ]

env:
  SEMVERBOT_VERSION: "1.0.0"

jobs:
  build:
    name: pipeline
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
        - 
      - name: set up path
        run: |
          mkdir bin
          echo "$(pwd)/bin" >> $GITHUB_PATH

      - name: install semverbot
        run: |
          curl -o bin/sbot -L https://github.com/restechnica/semverbot/releases/download/v$SEMVERBOT_VERSION/sbot-linux-amd64
          chmod +x bin/sbot
          
      - name: update version
        run: |
          sbot update version
          echo "CURRENT_VERSION=$(sbot get version)" >> $GITHUB_ENV
          echo "RELEASE_VERSION=$(sbot predict version)" >> $GITHUB_ENV
          echo "current version: $CURRENT_VERSION"
          echo "next version: $RELEASE_VERSION"
          
      ... build / publish ...
          
      - name: release version
        run: |
          sbot release version
          sbot push version
```

### Development workflow

A typical development workflow when working with `sbot`:

`[create branch 'feature/my-feature' from main/master]` > `[make changes]` > `[push changes]` > `[create pull request]` > `[approve pull request]` > `[merge pull request]` > `[trigger pipeline]` >
`[calculate next version based on branch name]` > `[build application]` > `[publish artifact]` > `[semverbot release & push version]`

## Why Semverbot?

There are several reasons why you should consider using `sbot` for your semver versioning.

`sbot` is originally made for large scale IT departments which maintain hundreds, if not thousands, of code repositories.
Manual releases for each of those components and their subcomponents cost a considerable amount of developer time.

1. Standardize how your releases are tagged
2. Automate the releasing process for potentially thousands of code repositories

### Automation and pipelines

`sbot` automates the process of tagging releases for you.
* `sbot` uses `git` under the hood, which is today's widely adopted version control system
* `sbot` does **not** use a file to keep track of the version
  * no pipeline loops
  * no need to maintain the version in two places, e.g., both a package.json file and git tags
* `sbot` is ready to be used in pipelines out of the box

Note: it is still possible to use `sbot` and file-based versioning tools side-by-side

### Convenience

* `sbot` is designed to be used by both developers and pipelines
* `sbot` is platform independent
  * support for Windows, Linux and macOS
  * no dependency on 'complex' `npm`, `pip` or other package management installations
* `sbot` heavily simplifies incrementing semver levels based on git information
  * today's `git` projects already hold a lot of useful semver information, e.g., branch names like `feature/xxx` or commit messages like `[fix] xxx`
  * no need to create and maintain custom code for semver level detection

### Configurability

* `sbot` supports a well-documented configuration file
  * intuitively customize how patch, minor and major levels are detected
* `sbot` supports several flags to override parts of the configuration file on the fly
