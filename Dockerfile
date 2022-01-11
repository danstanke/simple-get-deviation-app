FROM golang:1.17.5-alpine
COPY ./server/. /go/src/.
WORKDIR /go/src
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/server ./
CMD ["./server"]