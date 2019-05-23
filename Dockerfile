FROM golang:1.10
COPY . /go/src/github.com/pyxida/hello-world-k8s
WORKDIR /go/src/github.com/pyxida/hello-world-k8s
RUN go get github.com/prometheus/client_golang/prometheus
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o hello-world-k8s

FROM scratch
COPY --from=0 /go/src/github.com/pyxida/hello-world-k8s /
EXPOSE 8080
ENTRYPOINT ["/hello-world-k8s"]