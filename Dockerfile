FROM ubuntu:20.04
USER root
ARG GO_VERSION=1.16.6
ARG GO_PACKAGE=go${GO_VERSION}.linux-amd64.tar.gz

ENV GOROOT /opt/go
ENV GOPATH /opt/gopkg
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /build/kansible

RUN    apt update && apt install wget curl ansible nano sshpass  rsync -y  \
    && wget https://dl.google.com/go/${GO_PACKAGE} \
    && tar -zxvf ${GO_PACKAGE} -C /opt \
    && apt install ansible -y

COPY . .
RUN    /opt/go/bin/go mod download \
    && /opt/go/bin/go build -o server/kansible server/main.go

RUN mkdir /root/.ssh  \
    && touch /root/.ssh/config \
    && echo "Host *" >> /root/.ssh/config \
    && echo "\tStrictHostKeyChecking no" >> /root/.ssh/config \
    && echo "\tUserKnownHostsFile /dev/null" >> /root/.ssh/config \
    && ssh-keygen -t rsa -q -P "" -f ~/.ssh/id_rsa


COPY ansible.cfg /etc/ansible/ansible.cfg

RUN echo 'server/kansible' >> entrypoint.sh
EXPOSE 8080
CMD ["sh","entrypoint.sh"]
