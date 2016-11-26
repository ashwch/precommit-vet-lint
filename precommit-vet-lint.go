package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	lintErrors := make([]string, 0)
	vetErrors := make([]string, 0)
	stagedFiles := GetStagedFiles()

	temp := TempDir{}
	temp.Create()
	defer temp.Close()

	for _, file := range stagedFiles {
		fileDir, fileName := filepath.Split(file)
		tempFileDir := filepath.Join(temp.Path, fileDir)
		if _, err := os.Stat(tempFileDir); os.IsNotExist(err) {
			err = os.MkdirAll(tempFileDir, os.ModePerm)
		}
		tempFilePath := filepath.Join(tempFileDir, fileName)
		stagedContent := GetStagedContent(file)
		err := ioutil.WriteFile(tempFilePath, stagedContent, 0655)
		CheckError(err)
		lintErrorMessage, err := GetLintErrors(tempFilePath)
		if err != nil {
			lintErrors = append(lintErrors, strings.Replace(lintErrorMessage, tempFilePath, file, -1))
		}

		vetErrorMessage, err := GetVetErrors(tempFilePath)
		if err != nil {
			vetErrors = append(vetErrors, strings.Replace(vetErrorMessage, tempFilePath, file, -1))
		}
	}
	if len(lintErrors) != 0 {
		fmt.Println("Following linting error(s) found:")
		fmt.Println()
		for _, errMsg := range lintErrors {
			fmt.Println(errMsg)
		}
		fmt.Println()
		fmt.Println()
	}

	if len(vetErrors) != 0 {
		fmt.Println("Following vet error(s) found:")
		fmt.Println()
		for _, errMsg := range vetErrors {
			fmt.Println(errMsg)
		}
	}
	if len(lintErrors) != 0 || len(vetErrors) != 0 {
		os.Exit(1)
	}
}
