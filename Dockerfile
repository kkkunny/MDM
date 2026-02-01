FROM golang:1.25.6-alpine3.22 AS builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /mdm
COPY . .
RUN go build --tags release -o output/mdm .


FROM cnk3x/xunlei:v3.20.2 AS final

WORKDIR /mdm
COPY --from=builder /mdm/output/* /mdm
EXPOSE 80
ENTRYPOINT ["/mdm/mdm"]