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

## Developing

This repository contains scripts that run in linux or mac for installing pre-commit hooks. The files read from release. You must attach binaries if you add something new.
