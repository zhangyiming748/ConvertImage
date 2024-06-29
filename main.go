package main

import (
	"github.com/zhangyiming748/ConvertImage/constant"
	"github.com/zhangyiming748/ConvertImage/conv"
	"github.com/zhangyiming748/ConvertImage/mediainfo"
	"github.com/zhangyiming748/ConvertImage/util"
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	setLog()
}
func main() {
	if root := os.Getenv("root"); root == "" {
		log.Printf("$root为空,使用默认值:%v\n", constant.GetRoot())
	} else {
		constant.SetRoot(root)
		log.Printf("$root不为空:%v\n", constant.GetRoot())
	}
	// TODO  容器中不需要使用控制台方法退出
	// go util.ExitAfterRun()

	err := filepath.Walk(constant.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			log.Printf("准备处理的文件夹%v\n", info.Name())

			files := util.GetAllFiles(absPath)
			for _, file := range files {
				conv.ProcessImage(*mediainfo.GetBasicInfo(file))
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Error:", err)
	}

	files := util.GetAllFiles(constant.Root)

	for _, file := range files {
		conv.ProcessImage(*mediainfo.GetBasicInfo(file))
	}
}

func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "mylog.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
