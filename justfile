run: test cover

buf:
     buf generate --template proto/buf.go.yaml

check:
    go vet ./...
    go fmt ./...

test:
    go test ./...

cover:
    go test -coverprofile=coverage.out ./...
    go tool cover -func coverage.out