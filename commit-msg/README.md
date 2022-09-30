## Git Conventional Commits

## NOTE YOU MUST HAVE

```bash
npm install --global git-conventional-commits
```

## Configuration you must have `git-conventional-commits.json` file with folows:

```json
{
    "convention" : {
      "commitTypes": [
        "feat",
        "fix",
        "perf",
        "refactor",
        "style",
        "test",
        "build",
        "ops",
        "docs",
        "merge"
      ],
      "commitScopes": [],
      "releaseTagGlobPattern":  "v[0-9]*.[0-9]*.[0-9]*",
      "issueRegexPattern": "(^|\\s)#\\d+(\\s|$)"
    },

    "changelog" : {
      "commitTypes": [
        "feat",
        "fix",
        "perf",
        "merge"
      ],
      "includeInvalidCommits": true,
      "commitScopes": [],
      "commitIgnoreRegexPattern": "^WIP ",
      "headlines": {
        "feat": "Features",
        "fix": "Bug Fixes",
        "perf": "Performance Improvements",
        "merge": "Merged Branches",
        "breakingChange": "BREAKING CHANGES"
      },

      "commitUrl": "https://github.com/ACCOUNT/REPOSITORY/commit/%commit%",
      "commitRangeUrl": "https://github.com/ACCOUNT/REPOSITORY/compare/%from%...%to%?diff=split",
      "issueUrl": "https://github.com/ACCOUNT/REPOSITORY/issues/%issue%"
    }
  }
```

## also must have `commit-msg` file in `githokks/` folder with folows:

```bash
#!/bin/sh

PATH="/c/Program Files/nodejs:$HOME/AppData/Roaming/npm/:$PATH"
git-conventional-commits commit-msg-hook "$1"

```


## NOTE Config for GitHooks

git config core.hooksPath githooks


