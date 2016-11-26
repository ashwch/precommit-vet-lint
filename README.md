Precommit-vet-lint
===

A precommit hook that runs on your staged Go file and reports errors returned by [Golint](https://github.com/golang/lint) and [Vet](https://golang.org/cmd/vet/).

### Installation
---

```bash
go get -u github.com/golang/lint/golint
go get -u github.com/ashwch/precommit-vet-lint
```

### Usage
---

Copy the contents of pre-commit file included with this repo in your `.git/hooks/pre-commit` file. Make sure the file as executable permissions.


### Example

```bash
$ git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

    new file:   helpers.go
    new file:   precommit-vet-lint.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)

    precommit-vet-lint

$ git commit -m "Added core files"
Running Golint and Vet...
Failed Following linting error(s) found:
helpers.go:14:1: exported function CheckError should have comment or be unexported
helpers.go:20:1: exported function GetStagedFiles should have comment or be unexported
helpers.go:48:1: exported function GetProjectBasePath should have comment or be unexported
helpers.go:53:1: exported function GetStagedContent should have comment or be unexported
helpers.go:60:1: exported function GetLintErrors should have comment or be unexported
helpers.go:69:1: exported function GetVetErrors should have comment or be unexported
helpers.go:83:6: exported type TempDir should have comment or be unexported
helpers.go:87:1: exported method TempDir.Create should have comment or be unexported
helpers.go:93:1: exported method TempDir.Close should have comment or be unexported.
```
