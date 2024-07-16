package main

import (
	"github.com/zhangyiming748/ConvertImage/constant"
	"github.com/zhangyiming748/ConvertImage/conv"
	mylog "github.com/zhangyiming748/ConvertImage/log"
	"github.com/zhangyiming748/ConvertImage/util"
	"log"
	"os"
	"path/filepath"
)

func init() {
	mylog.SetLog()
}
func main() {
	if root := os.Getenv("root"); root == "" {
		log.Printf("$root为空,使用默认值:%v\n", constant.GetRoot())
	} else {
		constant.SetRoot(root)
		log.Printf("$root不为空:%v\n", constant.GetRoot())
	}
	err := filepath.Walk(constant.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			log.Printf("准备处理的文件夹:%v\n", info.Name())
			files := util.GetAllFiles(absPath)
			for _, file := range files {
				conv.ProcessImage(*util.GetBasicInfo(file))
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Error:", err)
	}
}
