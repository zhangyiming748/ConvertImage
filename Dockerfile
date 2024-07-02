FROM golang:1.22.4-bookworm
LABEL authors="zen"
COPY debian.sources /etc/apt/sources.list.d/
RUN apt update
RUN apt install -y dos2unix ffmpeg
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN go mod vendor
ENTRYPOINT ["go","run","/root/app/main.go"]