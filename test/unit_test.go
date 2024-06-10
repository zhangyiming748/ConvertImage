package test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

/*
递归方法获得全部层级的子文件夹
*/
func TestGetAllFolder(t *testing.T) {
	path := "/mnt/c/Users/zen" // 请替换为你需要的文件夹路径
	getAllFolder(path)
}

func getAllFolder(path string) {
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			log.Println(absPath)
		}
		return nil
	})
	if err != nil {
		log.Println("Error:", err)
	}
}
