
install:
	go install cmd/conmake.go

build:
	go build cmd/conmake.go

test:
	go test ./pkg/...

coverage:
	go test -coverprofile cover.out ./pkg/...
	go tool cover -func cover.out
	rm cover.out