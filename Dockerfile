FROM golang:1.25.6-alpine3.22 AS builder

ENV GOPROXY=https://goproxy.cn,direct
RUN apk --no-cache add build-base
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 go build --tags release -o output/mdm .


FROM alpine:3.22 AS final
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++ && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/output/* /app
COPY --from=builder /app/static /app/static

EXPOSE 80

ENTRYPOINT ["/app/mdm"]