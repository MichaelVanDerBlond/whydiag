APP=whydiag

build:
	go build -o bin/$(APP) ./cmd/whydiag

test:
	go test ./...

run:
	go run ./cmd/whydiag

fmt:
	go fmt ./...

clean:
	rm -rf bin
