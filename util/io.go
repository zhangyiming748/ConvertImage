package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		log.Printf("按行读文件出错:%v\n", err)
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}

// 按行写文件
func WriteByLine(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
	return

}

/*
获取当前文件夹和全部子文件夹下视频文件
*/

func GetAllFiles(root string) (files []string) {
	patterns := []string{"jpeg", "jpg", "png", "webp", "tif"}
	for _, pattern := range patterns {
		files = append(files, getFilesByExtension(root, pattern)...)
	}
	return files
}

/*
获取当前文件夹和全部子文件夹下指定扩展名的全部文件
*/
func getFilesByExtension(root, extension string) []string {
	var files []string
	defer func() {
		if err := recover(); err != nil {
			log.Println("获取文件出错")
			os.Exit(-1)
		}
	}()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
