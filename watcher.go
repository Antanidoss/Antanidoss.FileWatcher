package fileWatcher

type Watcher struct {
	DirectoryPath          string
	WatchNestedDirectories bool
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

func (watcher Watcher) Start() {

}

func (watcher Watcher) Stop() {

}

func NewWatcher(directoryPath string, watchNestedDirectories bool, onCreatedFile func(message EventFileWatcherMessage), onRemoveFile func(message EventFileWatcherMessage), onAnyChange func(message EventFileWatcherMessage)) (*Watcher, error) {
	if onCreatedFile == nil && onRemoveFile == nil && onAnyChange == nil {
		panic("All event handlers are equal nil")
	}

	return &Watcher{
		DirectoryPath:          directoryPath,
		WatchNestedDirectories: watchNestedDirectories,
		OnCreatedFile:          onCreatedFile,
		OnRemoveFile:           onRemoveFile,
		OnAnyChange:            onAnyChange}, nil
}
