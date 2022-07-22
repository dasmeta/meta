## What we have.

We run our hooks on every commit to automatically point out issues in code such as terraform fmt, terraform docs, trailing whitespace, detect-aws-credentials, check-merge-conflict, detect-private-key.

## INSTALL in Linux 

Run :
  ```bash
  curl -fsSL -O https://github.com/dasmeta/meta/releases/download/v0.1.0/add-this-linux.sh 
  ./add-this-linux.sh
  ```

and it will do all automanically like` create needed folders and files and do needed commands (install,run,etc.)

## NOTE Config for GitHooks

git config core.hooksPath githooks

## INSTALL in MacOS

Run :
  ```bash
  curl -fsSL -O https://github.com/dasmeta/meta/releases/download/v0.1.0/add-this-mac.sh 
  ./add-this-mac.sh
  ```

## If you want to do it all manually for use pre-commit githooks you need to have installed pre-commit see here how to install.

<!-- markdownlint-disable no-inline-html -->

* [`pre-commit`](https://pre-commit.com/#install)  

<!-- markdownlint-enable no-inline-html -->

### You need to have [`.pre-commit-config.yaml`](./.pre-commit-config.yaml) file in the top of your repo,
the hooks you enable you need to set in this file, for example to enable "end-of-file-fixer" and "detect-aws-credentials" hooks you can set this config

```bash
repos:
- repo: https://github.com/pre-commit/pre-commit-hooks
  rev: v4.2.0
  hooks:
    - id: end-of-file-fixer
    - id: detect-aws-credentials
```

### Install the pre-commit hook globally run this commands.

```bash
DIR=~/.git-template
git config --global init.templateDir ${DIR}
pre-commit init-templatedir -t pre-commit ${DIR}
```

### Install pre-commit hook PATH in your folder.

```bash
git config core.hooksPath {PATH-TO-HOOKS}

example

git config core.hooksPath githooks