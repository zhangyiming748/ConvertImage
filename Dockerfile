# 第一阶段：构建amd64镜像
FROM golang:1.22.4-bookworm as builder-amd64
LABEL authors="zen"
COPY debian.sources /etc/apt/sources.list.d/
RUN apt update
RUN apt install -y dos2unix
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN dos2unix /root/app/install-retry.sh
RUN chmod +x /root/app/install-retry.sh
RUN /root/app/install-retry.sh ffmpeg
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv

# 第二阶段：构建arm64镜像
FROM arm64/golang:1.22.4-bookworm as builder-arm64
LABEL authors="zen"
COPY debian.sources /etc/apt/sources.list.d/
RUN apt update
RUN apt install -y dos2unix
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN dos2unix /root/app/install-retry.sh
RUN chmod +x /root/app/install-retry.sh
RUN /root/app/install-retry.sh ffmpeg
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv

# 第三阶段：合并amd64和arm64镜像
FROM scratch
COPY --from=builder-amd64 /usr/local/bin/conv /usr/local/bin/conv
COPY --from=builder-arm64 /usr/local/bin/conv /usr/local/bin/conv
CMD ["conv"]
