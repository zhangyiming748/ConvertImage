#!/usr/bin/env bash
echo 删除旧的日志文件
find . -type f -name "*.log" -exec rm {} \;
echo 格式化当前目录下go文件
find . ! -path "./vendor/*" -name "*.go" -exec golog -w {} \;
echo 删除多余隐藏文件
find . -name "*DS_Store*" -exec rm {} \;
echo 编译二进制文件forLinux32
# CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o process4Linux32 main.go
echo 编译二进制文件forLinux64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o process4Linux64 main.go
echo 编译二进制文件forRaspi
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o process4Raspi main.go
echo 编译二进制文件forRaspi64
# CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o process4Raspi64 main.go
echo 编译二进制文件forWin32.exe
# CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o process4Win32.exe main.go
echo 编译二进制文件forWin64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o process4Win64.exe main.go
echo 编译二进制文件forMac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o process4Mac main.go
echo 编译二进制文件forM1
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o process4M1 main.go
echo 编译二进制文件forAndroid
CGO_ENABLED=0 GOOS=android GOARCH=arm64 go build -o process4Android main.go
echo 统计代码行数
find . -type f  ! -path "*vendor*" -name "*.go" -print0 | xargs -0 wc -l | tee lines.txt
echo '转换\r\n为\n'
# shellcheck disable=SC2038
find . -type f ! -path "./.git/*" ! -path "./.idea/*" ! -name "build.sh" ! -name "conf.ini" -print0 | xargs -0 dos2unix
# find . -name "*.go" -exec dos2unix {} \;