install:
	go mod download
run:
	./scripts/run_dev.sh
test:
	go test -v ./...
vet:
	go vet cmd/main.go
vet_shadow:
	go vet -vettool=$(which shadow) cmd/main.go
build:
	go build -ldflags="-X 'main.Version=v1.0.0'" -o linktly cmd/main.go
staticcheck:
	staticcheck ./...
audit:
	govulncheck -mode binary -show verbose linktly
gosec:
	gosec ./...
