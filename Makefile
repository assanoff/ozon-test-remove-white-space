.PHONY: test
test:
	go test -v -timeout 3s ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem  -timeout 3s ./...