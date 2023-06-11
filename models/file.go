package models

import "time"

type File struct {
	Name         string
	FullName     string
	ModTime      time.Time
	CreationTime time.Time
}
