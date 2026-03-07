FROM golang:1.25.6-alpine3.22 AS builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build --tags release -o output/mdm .


FROM alpine:3.22 AS final

RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk update && apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone \
ENV TZ Asia/Shanghai
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app
COPY --from=builder /app/output/* /app
COPY --from=builder /app/static /app/static

EXPOSE 80

ENTRYPOINT ["/app/mdm"]