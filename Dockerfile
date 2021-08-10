FROM ubuntu:20.04

ARG GO_VERSION=1.16.6
ARG GO_PACKAGE=go${GO_VERSION}.linux-amd64.tar.gz

ENV GOROOT  /opt/go
ENV GOPATH /opt/gopkg
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /build/kansible

RUN    apt update && apt install wget curl ansible -y  \
    && wget https://dl.google.com/go/${GO_PACKAGE} \
    && tar -zxvf ${GO_PACKAGE} -C /opt \
    && apt install ansible -y

COPY . .
RUN    go mod download \
    && go build -o server/kansible server/main.go

RUN echo 'server/kansible' >> entrypoint.sh
EXPOSE 8080
CMD ["sh","entrypoint.sh"]
