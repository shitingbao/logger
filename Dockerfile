FROM golang AS logtage

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn

WORKDIR /logger

COPY app/ app/
COPY lib/ lib/
COPY go.mod go.mod

WORKDIR /logger/app/cmd

RUN go build

FROM ubuntu
#重新构建，减少体积，这里只需要编译生成的可执行文件，配置文件，前端dist文件即可
WORKDIR /opt
COPY --from=logtage  /logger/app/cmd/cmd .
COPY --from=logtage  /logger/app/cmd/config.toml .
#no-install-recommends参数来避免安装非必须的文件，从而减小镜像的体积

EXPOSE 5000
EXPOSE 8080

ENTRYPOINT ["/opt/cmd"]