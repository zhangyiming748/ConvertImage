package main

import (
	"fmt"
	"github.com/zhangyiming748/ConvertImage/constant"
	"github.com/zhangyiming748/ConvertImage/conv"
	"github.com/zhangyiming748/ConvertImage/mediainfo"
	"github.com/zhangyiming748/ConvertImage/sql"
	"github.com/zhangyiming748/ConvertImage/util"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if root := os.Getenv("root"); root == "" {
		slog.Info("$root为空,使用默认值", slog.String("$root", constant.GetRoot()))
	} else {
		constant.SetRoot(root)
		slog.Info("$root不为空", slog.String("$root", root))
	}
	// TODO  容器中不需要使用控制台方法退出
	// go util.ExitAfterRun()
	if level := os.Getenv("level"); level == "" {
		slog.Info("$level为空,使用默认值", slog.String("$level", constant.GetLevel()))
		setLog(constant.GetLevel())
	} else {
		constant.SetLevel(level)
		slog.Info("$level不为空", slog.String("$level", level))
		setLog(constant.GetLevel())
	}
	sql.SetEngine()

	err := filepath.Walk(constant.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			fmt.Printf("准备处理的文件夹%v\n", info.Name())
			files := util.GetAllFiles(absPath)
			for _, file := range files {
				conv.ProcessImage(*mediainfo.GetBasicInfo(file))
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	files := util.GetAllFiles(constant.Root)

	for _, file := range files {
		conv.ProcessImage(*mediainfo.GetBasicInfo(file))
	}
}
func setLog(level string) {
	var opt slog.HandlerOptions
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	fp := strings.Join([]string{constant.GetRoot(), "ConvImage.log"}, string(os.PathSeparator))
	fmt.Printf("数据库位置%v\n", fp)
	logf, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
