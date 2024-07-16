package constant

import "runtime"

var (
	Root    string = "/data"          // 工作目录 如果为空  默认/data
	CpuNums int    = runtime.NumCPU() // 核心数
)

const (
	MaxCPU = 12
)

func GetCpuNums() int {
	return CpuNums
}

func GetRoot() string {
	return Root
}
func SetRoot(s string) {
	Root = s
}
