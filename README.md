
# Semverbot

[![github.com release badge](https://img.shields.io/github/release/restechnica/semverbot.svg)](https://github.com/restechnica/anyreleaser/semverbot/)
[![github.com workflow badge](https://github.com/restechnica/semverbot/workflows/main/badge.svg)](https://github.com/restechnica/semverbot/actions?query=workflow%3Amain)
[![go.pkg.dev badge](https://pkg.go.dev/badge/github.com/restechnica/semverbot)](https://pkg.go.dev/github.com/restechnica/semverbot)
[![goreportcard.com badge](https://goreportcard.com/badge/github.com/restechnica/semverbot)](https://goreportcard.com/report/github.com/restechnica/semverbot)
[![img.shields.io MPL2 license badge](https://img.shields.io/github/license/restechnica/semverbot)](./LICENSE)

A CLI which automates semver versioning based on git information.

## Why Semverbot?

There are several reasons why you should consider using `sbot` for your semver versioning.

### Automation and pipelines

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
* `sbot` is fast
* `sbot` heavily simplifies incrementing semver levels based on git information
  * today's `git` projects already hold a lot of useful semver information, e.g., branch names like `feature/xxx` or commit messages like `[fix] xxx`
  * no need to create and maintain custom code for semver level detection

### Configurability

* `sbot` supports a well-documented configuration file
  * intuitively customize how patch, minor and major levels are detected
* `sbot` supports some flags to override parts of the configuration file on the fly

## Requirements

`sbot` requires a `git` installation.

## How to install

`sbot` can be retrieved from GitHub or a Homebrew tap. Run `sbot -h` to validate the installation.
The tool is available for Windows, Linux and macOS.

### github

The following example works for a GitHub Workflow, other CI/CD tooling will require a different path setup.
The curl command remains the same.

```shell
SEMVERBOT_VERSION=0.1.2
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

## Commands

Each command has a `-h, --help` flag available.

### `sbot get version`

Gets the current version, which is the latest `git` annotated tag without any prefix.

### `sbot init`

Generates a configuration with defaults, see [configuration defaults](#defaults).

### `sbot predict version [-m, --mode] <mode>`

Gets the next version, without any prefix. Uses a mode to detect which semver level it should increment. Defaults to mode `auto`.
See [Modes](#modes) for more documentation on the supported modes.

### `sbot push version`

Pushes the latest `git` tag. Equivalent to `git push origin {prefix}{version}`.

### `sbot release version [-m, --mode] <mode>`

Creates a new version, which is a `git` annotated tag. Uses a mode to detect which semver level it should increment.
Defaults to mode `auto`. See [Modes](#modes) for more documentation on the supported modes.

### `sbot update version`

Fetches all tags with `git` to make sure the git repo has the latest tags available. Equivalent to `git fetch --unshallow`.
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

The branch name is matched against the ['semver.detection' configuration](#semverdetectionlevel).

### git-commit

Detects which semver level to increment based the **message** of the latest `git` commit.

The commit message is matched against the ['semver.detection' configuration](#semverdetectionlevel).

### major

Increments the `major` level.

### minor

Increments the `minor` level.

### patch

Increments the `patch` level.

## How to configure

`sbot` supports a configuration file. It looks for a `.semverbot.toml` file in the current working directory by default. `.json` and `.yaml` formats are also supported, but `.toml` is highly recommended.

### Defaults

`sbot init` generates the following configuration:

```toml
[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "v"

[semver]
mode = "auto"

[semver.detection]
patch = ["fix/", "[fix]"]
minor = ["feature/", "[feature]"]
major = ["release/", "[release]"]
```

## Configuration properties

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
The 'v' prefix, e.g. `v1.0.1` is used by default due to its popularity, e.g. some Golang tools completely depend on it.

Note: `sbot` will always display the version without the prefix.

### semver.detection

Some `sbot` semver modes require input to detect which semver level should be incremented.
Each level will be assigned a collection of matching strings, which are to be matched against `git` information.

See [Modes](#modes) for documentation about the supported modes.

### semver.detection.[level]

An array of strings which are to be matched against git information, like branch names and commit messages.
Whenever a match happens, `sbot` will increment the corresponding level.

### semver.mode
`sbot` supports multiple modes to detect which semver level it should increment. Each mode works with different criteria.
A `mode` flag enables you to switch modes on the fly.

See [Modes](#modes) for documentation about the supported modes.

## Examples

## Local

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
SEMVERBOT_VERSION=0.1.2
mkdir bin
echo "$(pwd)/bin" >> $GITHUB_PATH
curl -o bin/sbot -L https://github.com/restechnica/semverbot/releases/download/v$SEMVERBOT_VERSION/sbot-linux-amd64
chmod +x bin/sbot

# preparation
sbot update version
echo "RELEASE_VERSION=$(sbot predict version)" >> $GITHUB_ENV

# usage
echo "current version: $(sbot get version)"
echo "next version: $RELEASE_VERSION"
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
  SEMVERBOT_VERSION: "0.1.2"

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
          
      - name: prepare release
        run: |
          sbot update version
          echo "RELEASE_VERSION=$(sbot predict version)" >> $GITHUB_ENV    
          
      - name: release
        run: |
          echo "current version: $(sbot get version)"
          echo "next version: $RELEASE_VERSION"

          sbot release version
          sbot push version
```
