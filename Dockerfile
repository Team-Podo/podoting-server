#!/bin/bash

FROM amd64/golang:1.19 as build

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . .

RUN go mod tidy

RUN go build -o main main.go

# stage 2: Go 애플리케이션 바이너리 파일과 필요한 파일만 alpine 컨테이너에 복사
FROM alpine:latest
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/.env .
COPY --from=build /app/firebase-sdk.json .

EXPOSE 8080

# 컨테이너가 시작할 때 파일 실행
CMD ["/app/main"]
