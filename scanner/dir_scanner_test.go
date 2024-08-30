package scanner

import (
	"testing"
	"testing/fstest"
)

func TestFolderScan(t *testing.T) {
	mockFS := fstest.MapFS{
		"folder_1/folder_2/folder_3/file_1.json": {Data: []byte("testing data 1")},
		"folder_1/folder_2/folder_3/file_2.json": {Data: []byte("testing data 2")},
		"folder_1/folder_2/folder_3/file_3.json": {Data: []byte("testing data 3")},
		"folder_1/folder_2/folder_3/file_4.json": {Data: []byte("testing data 4")},
		"folder_1/folder_2/folder_3/file_5.json": {Data: []byte("testing data 5")},
		"folder_1/folder_2/folder_3/file_6.json": {Data: []byte("testing data 6")},
	}

	path := "folder_1/folder_2/folder_3"
	want := len(mockFS)

	output, err := ScanPath(mockFS, path)
	if err != nil {
		t.Fatal(err)
	}

	got := len(output)

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}
