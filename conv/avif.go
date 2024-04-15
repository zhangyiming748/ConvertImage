package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertImage/mediainfo"
	"github.com/zhangyiming748/ConvertImage/replace"
	"github.com/zhangyiming748/ConvertImage/sql"
	"github.com/zhangyiming748/ConvertImage/util"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

/*
转换一张图片为AVIF
*/
func ProcessImage(in mediainfo.BasicInfo) {
	c := new(sql.Conv)
	defer c.SetOne()
	safeDelete := false
	defer func() {
		if err := recover(); err != nil {
			slog.Error("处理图片失败", slog.Any("错误", err))
		} else {
			//如果可以安全删除
			if safeDelete {
				slog.Info("处理图片成功", slog.String("文件名", in.FullPath))
				if err = os.Remove(in.FullPath); err != nil {
					slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
				} else {
					slog.Debug("删除成功", slog.Any("源文件", in.FullName))
				}
			}
		}
	}()
	cleanName := replace.ForFileName(in.PurgeName)
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), cleanName, ".avif"}, "")

	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-c:v", "libaom-av1", "-still-picture", "1", out)
	slog.Debug("ffmpeg", slog.Any("生成的命令", fmt.Sprint(cmd)))
	util.ExecCommand(cmd, "")
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	//todo 如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		safeDelete = true
	}

	c.Src = in.FullPath
	c.Dst = out
	c.SrcSize = originsize
	c.DstSize = aftersize
	if !safeDelete {
		c.IsBigger = true
	} else {
		c.IsBigger = false
	}
}
