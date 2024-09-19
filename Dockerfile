FROM golang:1.22-alpine3.20 AS builder
WORKDIR $GOPATH/src/github.com/wildanie12/region-api/
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags="-w -s" -o /go/bin/region-api ./main.go

FROM alpine:3.20.3
COPY --from=builder /go/bin/region-api /go/bin/region-api
WORKDIR /go/bin/
