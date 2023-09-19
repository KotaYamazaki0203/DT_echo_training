# Dockerfile
FROM golang:1.18.10-alpine

# ログに出力する時間をJSTにするため、タイムゾーンを設定
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

ENV ROOT=/go/src/app
WORKDIR ${ROOT}/src

# コーディングしないファイルはCOPYしローカルと同期しないようにする
# GoLandとコンテナの兼ね合いでgo.modはそれぞれの環境で中身が異なるため、特に同期させたくない
COPY ./app/go.mod ${ROOT}/go.mod

# ModuleモードをON
ENV GO111MODULE=on