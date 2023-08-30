## What we have.

We run our hooks on every commit to automatically point out issues in code such as terraform fmt, terraform docs, trailing whitespace, detect-aws-credentials, check-merge-conflict, detect-private-key.

## INSTALL in Linux

Run :
  ```bash
    bash <(wget -qO- https://github.com/dasmeta/meta/releases/download/v0.1.0/add-this-linux.sh)
  ```

and it will do all automanically like` create needed folders and files and do needed commands (install,run,etc.)

## NOTE Config for GitHooks

git config core.hooksPath githooks

## INSTALL in MacOS

Run :
  ```bash
    bash <(wget -qO- https://github.com/dasmeta/meta/releases/download/v0.1.0/add-this-mac.sh)
  ```

## USAGE

```bash
Meta is a command-line tool to generate semantic-release and pre-commit hooks in your projects.
It helps you create needed files and run it and much more.

Usage:
   meta {flags}
   meta <command> {flags}

Commands:
   help                          displays usage information
   pre-commit                    This command creates .pre-commit.yaml files and after running scripts for needed pre-commit checks and outputs that files in the project directory.
   semantic-release              Automatically do versioning and generate changelogs
   version                       displays version number

Flags:
   -h, --help                    displays usage information of the application or a command (default: false)
   -v, --version                 displays version number (default: false)
```

## Example

```bash
meta pre-commit
meta semantic-release
meta help
meta version
```

## Developing

This repository contains scripts that run in linux or mac for installing pre-commit hooks. The files read from release. You must attach binaries if you add something new.

## Docker
### Check Version
`make version`
`docker run -it dasmeta/meta version`

### Test create
`make docker-build debug-gitlab`

### Build and publish docker image
`make docker-publish`
