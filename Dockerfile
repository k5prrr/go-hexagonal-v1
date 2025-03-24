# example https://github.com/zeromicro/go-zero/blob/master/tools/goctl/Dockerfile
# https://habr.com/ru/companies/otus/articles/660301/
# https://habr.com/ru/articles/647255/

FROM golang:1.16-alpine as builder
WORKDIR /build
COPY go.mod . # go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
