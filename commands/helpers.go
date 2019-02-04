package commands

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const oneKB = 1 * 1024
const oneMB = 1 * oneKB * 1024
const oneGB = 1 * oneMB * 1024
const oneTB = 1 * oneGB * 1024

func helpersGetFoldersAndFiles(folder string) ([]string, []string, []error) {
	items, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, nil, []error{err}
	}

	folders := []string{}
	files := []string{}
	errorList := []error{}

	for _, item := range items {
		switch mode := item.Mode(); {
		case mode.IsDir():
			completeFolderPath := filepath.Join(folder, item.Name())
			folders = append(folders, completeFolderPath)

			childFolders, childFiles, childErrorList := helpersGetFoldersAndFiles(completeFolderPath)
			if err != nil {
				errorList = append(errorList, childErrorList[:]...)
			} else {
				folders = append(folders, childFolders[:]...)
				files = append(files, childFiles[:]...)
			}

		case mode.IsRegular():
			files = append(files, filepath.Join(folder, item.Name()))
		}
	}

	return folders, files, errorList
}

func helpersCopyFile(source, destination string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return 0, err
	}
	defer destinationFile.Close()

	bytesCopied, err := io.Copy(destinationFile, sourceFile)

	return bytesCopied, err
}
