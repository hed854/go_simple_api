FROM golang:alpine as builder
WORKDIR /go/src/app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN chmod +x app

FROM scratch as runner
COPY --from=builder /go/src/app .
EXPOSE 8080
CMD ["./app"]
