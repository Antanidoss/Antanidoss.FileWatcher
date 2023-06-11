package models

type Watcher struct {
	DirectoryPath          string
	WatchNestedDirectories bool
	Working                bool
	TimeoutInSeconds       uint
	TrackedFiles           []File
	OnCreatedFile          func(message EventFileWatcherMessage)
	OnRemoveFile           func(message EventFileWatcherMessage)
	OnAnyChange            func(message EventFileWatcherMessage)
}

type NotificationType int64

const (
	CreatedFile NotificationType = 0
	RemovedFile NotificationType = 1
)

type EventFileWatcherMessage struct {
	DirectoryPath    string
	FilePath         string
	NotificationType NotificationType
}
