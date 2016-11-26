Precommit-vet-lint
===

A precommit hook that runs on your staged Go file and reports errors returned by [Golint](https://github.com/golang/lint) and [Vet](https://golang.org/cmd/vet/).

### Installation
---

```bash
go get -u github.com/ashwch/pre-commit-vet-lint
```

### Usage
---

Copy the contents of pre-commit file included with this repo in your `.git/hooks/pre-commit` file. Make sure the file as executable permissions.