package util

import (
	"bufio"
	"github.com/zhangyiming748/filetype"
	"io"
	"log"
	"os"
	"path/filepath"
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

func isImageWithoutAVIF(file string) bool {
	// Open a buf descriptor
	buf, _ := os.Open(file)
	// We only have to pass the buf header = first 261 bytes
	head := make([]byte, 261)
	buf.Read(head)
	if filetype.IsImage(head) {
		if filetype.IsMIME(head, "image/avif") {
			log.Printf("%v已经是avif文件，跳过\n", file)
			return false
		} else {
			log.Printf("%v不是avif文件，转换\n", file)
			return true
		}
	}
	return false
}

/*
获取当前文件夹和全部子文件夹下指定扩展名的全部文件
*/
func GetAllFiles(path string) []string {
	var files []string
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if isImageWithoutAVIF(p) {
				files = append(files, p)
			}
		}
		return nil
	})
	return files
}
