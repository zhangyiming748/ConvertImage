package util

import (
	"testing"
)

/*
go test -v -run TestGetAllFiles
*/
func TestGetAllFiles(t *testing.T) {
	dir := "/mnt/d/images"
	files := GetAllFiles(dir)
	for _, file := range files {
		t.Log(file)
	}
}
