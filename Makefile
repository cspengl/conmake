
install:
	go install cmd/conmake.go

build:
	go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH cmd/conmake.go

version:
	go build -ldflags "-X github.com/cspengl/conmake/pkg/utils.Version=$(version)" \
	-gcflags=-trimpath=$(GOPATH) \
	-asmflags=-trimpath=$(GOPATH) cmd/conmake.go

test:
	go test ./pkg/...

coverage:
	go test -coverprofile cover.out ./pkg/...
	go tool cover -func cover.out
	rm cover.out