package filewatcher

import (
	"fmt"
	"os"
	"time"

	"github.com/Antanidoss/fileWatcher/models"
)

func Start(watcher *models.Watcher) error {
	if watcher.OnCreatedFile == nil && watcher.OnRemoveFile == nil && watcher.OnAnyChange == nil {
		return fmt.Errorf("All event handlers are equal nil")
	}

	if _, err := os.Stat(watcher.DirectoryPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	watcher.Working = true

	go Watch(watcher)

	return nil
}

func Stop(watcher *models.Watcher) {
	watcher.Working = false
}

func Watch(watcher *models.Watcher) {
	for {
		time.Sleep(time.Duration(watcher.TimeoutInSeconds))

		if !watcher.Working {
			return
		}
	}
}
