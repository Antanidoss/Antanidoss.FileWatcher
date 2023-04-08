package models

type Folder struct {
	Files   []File
	Folders []Folder
}
