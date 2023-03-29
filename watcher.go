package fileWatcher

type Watcher struct {
	directoryPath          string
	watchNestedDirectories bool
	onCreatedFile          func(message EventFileWatcherMessage)
	onRemoveFile           func(message EventFileWatcherMessage)
	onAnyChange            func(message EventFileWatcherMessage)
}

type NotificationType int64

const (
	CreatedFile NotificationType = 0
	RemovedFile NotificationType = 1
)

type EventFileWatcherMessage struct {
	directoryPath    string
	filePath         string
	notificationType NotificationType
}
