package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
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

	watcher.TrackedFiles = *getFiles(watcher.DirectoryPath, watcher.WatchNestedDirectories)
	watcher.Working = true

	go watch(watcher)

	return nil
}

func Stop(watcher *models.Watcher) {
	watcher.Working = false
}

func getFiles(directoryPath string, watchNestedDirectories bool) *[]models.File {
	var files []models.File

	filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() /*|| (!watchNestedDirectories && directoryPath != filepath.Base(path))*/ {
			return nil
		}

		file := models.File{Name: info.Name(), FullName: path, ModTime: info.ModTime()}

		if runtime.GOOS == "windows" {
			d := info.Sys().(*syscall.Win32FileAttributeData)
			file.CreationTime = time.Unix(0, d.CreationTime.Nanoseconds())
		} else {

		}

		files = append(files, file)

		return nil
	})

	return &files
}

func watch(watcher *models.Watcher) {
	for {
		time.Sleep(time.Duration(watcher.TimeoutInSeconds))

		if !watcher.Working {
			return
		}

		if !watcher.WatchNestedDirectories {
			files := *getFiles(watcher.DirectoryPath, watcher.WatchNestedDirectories)

			for _, file := range files {
				if !isCreatedFile(&watcher.TrackedFiles, &file) {
					watcher.TrackedFiles = append(watcher.TrackedFiles, file)
					watcher.OnCreatedFile(models.EventFileWatcherMessage{DirectoryPath: watcher.DirectoryPath, FilePath: file.FullName, NotificationType: models.CreatedFile})
				} else if isRemovedFiles(&watcher.TrackedFiles, &file) {
					watcher.TrackedFiles = removeTrackedFile(watcher.TrackedFiles, &file)
					watcher.OnRemoveFile(models.EventFileWatcherMessage{DirectoryPath: watcher.DirectoryPath, FilePath: file.FullName, NotificationType: models.RemovedFile})
				}
			}
		}
	}
}

func isCreatedFile(currentFiles *[]models.File, file *models.File) bool {
	for _, item := range *currentFiles {
		if item.FullName == file.FullName {
			return true
		}
	}

	return false
}

func isRemovedFiles(currentFiles *[]models.File, file *models.File) bool {
	for _, item := range *currentFiles {
		if item.FullName == file.FullName {
			return false
		}
	}

	return true
}

func removeTrackedFile(trackedFiles []models.File, file *models.File) []models.File {
	var removedFileIndex int

	for i := 0; i < len(trackedFiles); i++ {
		if trackedFiles[i].FullName == file.FullName {
			removedFileIndex = i
		}
	}

	trackedFiles[removedFileIndex] = trackedFiles[len(trackedFiles)-1]
	return trackedFiles[:len(trackedFiles)-1]
}
