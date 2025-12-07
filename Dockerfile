FROM golang AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
# ENV GO111MODULE=on
# ENV GOPROXY=https://goproxy.cn
RUN go build -ldflags="-w -s" -trimpath -o /app/piclist-server ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/piclist-server .
ENTRYPOINT ["./piclist-server"]