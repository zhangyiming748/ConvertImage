package conv

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
转换一张图片为AVIF
*/
func ProcessImage(fp string) {
	ext := filepath.Ext(fp)
	out := strings.Replace(fp, ext, ".avif", -1)
	cmd := exec.Command("ffmpeg", "-y", "-i", fp, "-c:v", "libaom-av1", "-still-picture", "1", out)
	log.Printf("ffmpeg生成的命令:%v\n", cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Fatalf("命令运行产生错误%v\n", err)
	} else {
		os.Remove(fp)
	}
}
