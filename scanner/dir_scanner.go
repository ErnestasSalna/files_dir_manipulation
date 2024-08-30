package scanner

import (
	"io/fs"
)

func ScanPath(fsys fs.FS, path string) ([]fs.DirEntry, error) {
	files, _ := fs.ReadDir(fsys, path)
	return files, nil
}
