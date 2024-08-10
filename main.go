package main

import (
	"github.com/schollz/progressbar/v3"
	"github.com/zhangyiming748/ConvertImage/constant"
	"github.com/zhangyiming748/ConvertImage/conv"
	mylog "github.com/zhangyiming748/ConvertImage/log"
	"github.com/zhangyiming748/ConvertImage/util"
	"log"
	"os"
)

func main() {
	//os.Setenv("root", "/home/zen/share/Videos")
	if root := os.Getenv("root"); root == "" {
		log.Printf("$root为空,使用默认值:%v\n", constant.GetRoot())
	} else {
		constant.SetRoot(root)
		log.Printf("$root不为空:%v\n", constant.GetRoot())
	}
	mylog.SetLog()
	files := util.GetAllFiles(constant.GetRoot())
	length := len(files)
	bar := progressbar.New(length)
	defer bar.Finish()
	for i, file := range files {
		bar.Set(i)
		conv.ProcessImage(file)
	}
}
