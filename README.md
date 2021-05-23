# Semverbot

A CLI which automates semver versioning.

# Configuration

## Git

### Config

`sbot` requires a `user.name` and `user.email` to be configured with `git`.
Existing `git` configs will not be overwritten.
Without these configs `sbot` might show unexpected behaviour and will not be able to push tags.

### Tags

#### prefix

Controls the prefix of git tags. The `v prefix`, e.g. `v1.0.1` is used by default due to its popularity.
Some Golang tools completely depend on it.

#### fetch

`sbot` requires all tags to be fetched for it to work properly. This is often overlooked in pipelines,
where most `git` operations are shallow by default for performance reasons.
This config is set to `false` by default since it might not be necessary for everyone.
A `fetch` flag is provided to enable it on a case-by-case basis.

## Semver

### Detection
Depending on the chosen `mode`, `sbot` will require some configuration to be able to detect the `semver level` to apply.
Each level will be assigned a collection of matching strings, which will be matched against `git` information.

e.g.: while the `git-branch` mode is selected, `sbot` will match the strings against the name of the branch
where a merge originated from to identify the `semver level`.

### Mode
`sbot` supports multiple modes to detect and increment `semver levels`.
Each mode uses different criteria for their `semver level` detection system.
A `mode` flag is provided as well.

#### Available modes
**auto** (default)

Attempts a series of modes int he following order:
1. `git-branch`
1. `git-commit` - if `git-branch` failed
1. `patch` - if `git-commit` failed

**git-branch**

Detects and increments a `semver level` based on the **name** of the `git` branch from where a merge commit originated from.
This only works when the old branch has not been deleted yet. The branch name is matched with the `semver.detection` configuration.

**git-commit**

Detects and increments a `semver level` based on the `git` commit message.
The commit message is matched with the `semver.detection` configuration.

**major**

Increments the `major` level.

**minor**

Increments the `minor` level.

**patch**

Increments the `patch` level.

## Defaults

The following configuration is generated by `sbot init`
```toml
[git]

[git.config]
email = "sbot@sbot.com"
name = "semverbot"

[git.tags]
prefix = "v"
fetch = false

[semver]
mode = "auto"

[semver.detection]
patch = ["fix/"]
minor = ["feature/"]
major = ["release/"]
```
