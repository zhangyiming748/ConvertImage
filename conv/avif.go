package conv

import (
	"github.com/zhangyiming748/ConvertImage/replace"
	"github.com/zhangyiming748/ConvertImage/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
转换一张图片为AVIF
*/
func ProcessImage(in util.BasicInfo) {
	cleanName := replace.ForFileName(in.PurgeName)
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), cleanName, ".avif"}, "")
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	width, _ := strconv.Atoi(mi.Image.Width)
	height, _ := strconv.Atoi(mi.Image.Height)
	crf := FastMediaInfo.GetCRF("avif", width, height)
	if crf == "" {
		crf = "31"
	}
	cmd := exec.Command("ffmpeg", "-y", "-i", in.FullPath, "-c:v", "libaom-av1", "-crf", crf, "-still-picture", "1", out)
	log.Printf("ffmpeg生成的命令:%v\n", cmd.String())
	if err := util.ExecCommand(cmd, in.FullPath); err != nil {
		log.Fatalf("ffmpeg命令%s运行中产生错误:%v\n", cmd.String(), err)
	}
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	log.Printf("savesize: %f MB\n", sub)
	//todo 如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		os.Remove(in.FullPath)
	} else {
		log.Printf("新文件:s比源文件:%s还大 不删除源文件", in.FullPath, out)
	}
}
