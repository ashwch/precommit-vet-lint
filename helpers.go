package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// CheckError is a helper function for handling errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetStagedFiles return a list of all Go files currently staged.
func GetStagedFiles() []string {
	cmd := exec.Command("bash", "-c", "git diff --diff-filter=ACMRTUXB --cached HEAD --name-only | egrep '[.]go$'")
	stdout, _ := cmd.Output()
	stdoutString := strings.TrimSpace(string(stdout))
	if stdoutString != "" {
		return strings.Split(stdoutString, "\n")
	}
	return make([]string, 0)
}

func getPathToDotGit() string {
	cmd := exec.Command("bash", "-c", "git rev-parse --git-dir")
	stdout, err := cmd.Output()
	CheckError(err)

	dotGitPath := strings.TrimSpace(string(stdout))
	cwd, err := os.Getwd()
	CheckError(err)

	if dotGitPath == ".git" {
		return path.Join(cwd, dotGitPath)
	} else if dotGitPath == "." {
		return cwd
	}
	return dotGitPath

}

// GetStagedContent returns the content of a staged file.
func GetStagedContent(file string) []byte {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("git show :%s", file))
	stdout, err := cmd.Output()
	CheckError(err)
	return stdout
}

// GetLintErrors retuns the error returned by `golint` command on a file.
func GetLintErrors(file string) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("golint -set_exit_status %s", file))
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return removeLastLine(string(stdout)), err
	}
	return "", nil
}

// GetVetErrors retuns the error returned by `go vet` command on a file.
func GetVetErrors(file string) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("go vet %s", file))
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return removeLastLine(string(stdout)), err
	}
	return "", nil
}

func removeLastLine(str string) string {
	newlineIndex := strings.LastIndex(strings.TrimSpace(str), "\n")
	return str[:newlineIndex]
}

// TempDir is a struct that is used to store path to a temp directory.
type TempDir struct {
	Path string
}

// Create creates a temp directory and assigns the path to this directory to TempDir.Path
func (dir *TempDir) Create() {
	path, err := ioutil.TempDir("", "pre-commit-vet-lint")
	CheckError(err)
	dir.Path = path
}

// Close deletes the temporary directory whose path was stored under TempDir.Path
func (dir *TempDir) Close() {
	err := os.RemoveAll(dir.Path)
	CheckError(err)
}
