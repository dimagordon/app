FROM amazonlinux:2.0.20190508

RUN yum -y update && yum -y install git aws-cli tar gzip unzip which gcc postgresql

ENV GO_VERSION 1.13
RUN curl -sSL https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV GOPATH /go
ENV GOCACHE /tmp/gocache
ENV GOPROXY=https://proxy.golang.org
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

ENV GOOSE_VERSION 2.6.4
RUN curl -sSL --output $GOPATH/bin/goose https://github.com/lonja/goose/releases/download/v$GOOSE_VERSION/goose-linux64 && chmod +x $GOPATH/bin/goose

RUN go get -u github.com/volatiletech/sqlboiler && \
    go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql &&\
    go get -u github.com/derekparker/delve/cmd/dlv
