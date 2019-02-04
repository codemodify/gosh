package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// ParallelCP -
type ParallelCP struct {
	command string
}

// NewParallelCP -
func NewParallelCP() Command {
	return &ParallelCP{
		command: "pcp",
	}
}

// CanHandle -
func (thisRef *ParallelCP) CanHandle(command string) bool {
	return strings.HasPrefix(command, thisRef.command)
}

// Execute -
func (thisRef *ParallelCP) Execute(command string) error {
	args := strings.Split(command, " ")

	if len(args) < 3 {
		return fmt.Errorf("USGAE: %s {source_file|source_folder} {destination_file|destination_folder}", thisRef.command)
	}

	source := args[1]
	destination := args[2]

	fileInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	var totalCountersSync sync.RWMutex
	var bytesCopied int64
	var errorsToReturn = []error{}
	var timeElapsed time.Duration

	var startTime = time.Now()

	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		folders, files, errorList := helpersGetFoldersAndFiles(source)
		if len(errorList) > 0 {
			errorsToReturn = append(errorsToReturn, errorList[:]...)
		} else {
			for _, folder := range folders {
				newFolderName := strings.Replace(folder, source, filepath.Join(destination, fileInfo.Name()), 1)
				os.MkdirAll(newFolderName, os.ModePerm)
			}

			wg := sync.WaitGroup{}
			for _, file := range files {
				newFileName := strings.Replace(file, source, filepath.Join(destination, fileInfo.Name()), 1)

				wg.Add(1)
				go func(theFile string, theDestination string) {
					localBytesCopied, localErrorToReturn := helpersCopyFile(theFile, theDestination)

					totalCountersSync.Lock()
					bytesCopied = bytesCopied + localBytesCopied
					if localErrorToReturn != nil {
						errorsToReturn = append(errorsToReturn, localErrorToReturn)
					}
					totalCountersSync.Unlock()

					wg.Done()
				}(file, newFileName)
			}
			wg.Wait()
		}

	case mode.IsRegular():
		localBytesCopied, localErrorToReturn := helpersCopyFile(source, destination)

		bytesCopied = bytesCopied + localBytesCopied
		errorsToReturn = append(errorsToReturn, localErrorToReturn)
	}

	timeElapsed = time.Since(startTime)
	fmt.Println(fmt.Sprintf(
		"Total: %d bytes (%d KB, %d MB, %d GB, %d TB) in %v",
		bytesCopied,
		bytesCopied/oneKB,
		bytesCopied/oneMB,
		bytesCopied/oneGB,
		bytesCopied/oneTB,
		timeElapsed,
	))

	if len(errorsToReturn) > 0 {
		errorsAsStrings := []string{}
		for _, err := range errorsToReturn {
			errorsAsStrings = append(errorsAsStrings, err.Error())
		}

		return fmt.Errorf("ERROR: %s", strings.Join(errorsAsStrings, "\n"))
	}

	return nil
}
