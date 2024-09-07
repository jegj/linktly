install:
	go mod download
run:
	./scripts/run_dev.sh
test:
	go test -v ./...
