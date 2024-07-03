package util

import (
	"log"
	"os"
	"os/exec"
)

/*
执行命令过程中可以循环打印消息
*/
func ExecCommand(c *exec.Cmd, msg string) (e error) {
	log.Printf("%v:开始执行命令:%v\n", msg, c.String())

	output, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("命令运行出现错误:%v\n", err)
	} else {
		log.Println(string(output))
	}
	if isExitLabel() {
		log.Printf("命令端获取到退出状态,命令结束后退出:%v\n", c.String())
		os.Exit(0)
	}
	return nil
}

/*
判断古希腊掌管退出信号的文件是否存在
*/
func isExitLabel() bool {
	filePath := "/exit"

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Println("古希腊掌管退出信号的文件不存在")
		return false
	} else {
		log.Println("古希腊掌管退出信号的文件存在")
		return true
	}
}
