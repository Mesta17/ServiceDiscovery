FROM golang:1.6-onbuild
CMD go get github.com/Mesta17/ServiceDiscovery
CMD go run main.go