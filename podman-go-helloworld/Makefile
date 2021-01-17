
GO=podman run -v .:/opt/app -w /opt/app --rm docker.io/golang:alpine

build:
	$(GO) go build
	podman build -t 3musketeersio/go-helloworld .

test:
	$(GO) go test

deploy:
	$(GO) echo "deploy target not implemented"
