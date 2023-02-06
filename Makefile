GOFMT_FILES?=$$(find . -name '*.go')

default: build

build:
	go build -o server ./cmd/coffeetalk

debug:
	go build -o server -gcflags="all=-N -l" ./cmd/$(TARGET)

release:
	$(shell export GOOS=linux export GOARCH=amd64; go build -o server ./cmd/$(TARGET))

coverage:
	@sh -c "'$(CURDIR)/scripts/coverage.sh'"

profile:
	go tool pprof http://localhost:8082/debug/pprof/allocs

tests:
	@sh -c "'$(CURDIR)/scripts/test.sh'"

tidy:
	go mod tidy

deps:
	go mod download

mocks:
	mockery --case snake --inpackage --all --dir pkg/

docker: release
	docker build -t $(TARGET):$(VERSION) .

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: default build debug release coverage profile tests tidy deps mocks docker fmt fmtcheck
