package conv

import (
	"github.com/zhangyiming748/ConvertImage/mediainfo"
	"github.com/zhangyiming748/ConvertImage/replace"
	"github.com/zhangyiming748/ConvertImage/util"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
转换一张图片为AVIF
*/
func ProcessImage(in mediainfo.BasicInfo) {
	cleanName := replace.ForFileName(in.PurgeName)
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), cleanName, ".avif"}, "")

	cmd := exec.Command("ffmpeg", "-y", "-i", in.FullPath, "-c:v", "libaom-av1", "-still-picture", "1", out)
	log.Printf("ffmpeg生成的命令:%v\n", cmd.String())
	if err := util.ExecCommand(cmd, in.FullPath); err != nil {
		os.Exit(-1)
	}
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	log.Printf("savesize: %f MB\n", sub)
	//todo 如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		os.Remove(in.FullPath)
	}
}
