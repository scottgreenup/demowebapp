# ==============================================================================
FROM golang:alpine as builder
ADD . /go/src/github.com/scottgreenup/demowebapp
RUN go build -o /go/bin/demowebapp github.com/scottgreenup/demowebapp

# ==============================================================================
FROM alpine:latest
COPY --from=builder /go/bin/demowebapp /bin/demowebapp
EXPOSE 8080
ENTRYPOINT /bin/demowebapp
