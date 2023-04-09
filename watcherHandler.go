package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"

	"github.com/Antanidoss/fileWatcher/models"
)

func Start(watcher *models.Watcher) error {
	if watcher.OnCreatedFile == nil && watcher.OnRemoveFile == nil && watcher.OnAnyChange == nil {
		return fmt.Errorf("All event handlers are equal nil")
	}

	if watcher.DirectoryPath == "" {
		return fmt.Errorf("DirectoryPath cannot be empty")
	}

	if _, err := os.Stat(watcher.DirectoryPath); err != nil {
		return err
	}

	if watcher.TimeoutInSeconds == 0 {
		watcher.TimeoutInSeconds = 200
	}

	watcher.Working = true

	watcher.State = *createFolderState(watcher.DirectoryPath, watcher.WatchNestedDirectories)

	go watch(watcher)

	return nil
}

func Stop(watcher *models.Watcher) {
	watcher.Working = false
}

func createFolderState(directoryPath string, watchNestedDirectories bool) *models.Folder {
	if watchNestedDirectories {
		return createFolderStateTree(directoryPath)
	}

	folderFiles, err := ioutil.ReadDir(directoryPath)

	if err != nil {
		panic(err)
	}

	folder := models.Folder{Files: make([]models.File, len(folderFiles))}

	for _, file := range folderFiles {
		folder.Files = append(folder.Files, models.File{Name: file.Name(), FullName: file.Name()})
	}

	return &folder
}

func createFolderStateTree(directoryPath string) *models.Folder {

}

func watch(watcher *models.Watcher) {
	for {
		time.Sleep(time.Duration(watcher.TimeoutInSeconds))

		if !watcher.Working {
			return
		}

		if !watcher.WatchNestedDirectories {
			files, _ := ioutil.ReadDir(watcher.DirectoryPath)

			for _, file := range files {
				checkCreatedFile(watcher, file)
			}
		}
	}
}

func checkCreatedFile(watcher *models.Watcher, file fs.FileInfo) {

}
