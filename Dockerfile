FROM golang:1.11-stretch AS builder
WORKDIR /go/src/github.com/pollosp/microproxy/
RUN go get -d -v github.com/labstack/echo && go get -d -v github.com/labstack/echo/middleware
RUN useradd -u 10001 scratchuser
COPY server.go .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app/
COPY --from=builder /go/src/github.com/pollosp/microproxy/microproxy .
COPY --from=builder /etc/passwd /etc/passwd
COPY static /app/static/
USER scratchuser
ENTRYPOINT ["./microproxy"]
