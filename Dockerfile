FROM golang:1.25.6-alpine3.22 AS builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build --tags release -o output/mdm .


FROM alpine:3.22 AS final

WORKDIR /app
COPY --from=builder /app/output/* /app
COPY --from=builder /app/static /app/static

EXPOSE 80

ENTRYPOINT ["/app/mdm"]