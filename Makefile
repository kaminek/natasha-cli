NAME := natasha-cli
EXECUTABLE := $(NAME)
PACKAGES ?= $(shell go list ./... | grep -v /vendor/ | grep -v /_tools/)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./_tools/*")
HEADERS_CONFIG ?= natasha_headers.yml

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf bin/ $(DIST)/
	rm -f pkg/headers/cgo_helpers.go pkg/headers/cgo_helpers.h pkg/headers/cgo_helpers.c
	rm -f pkg/headers/const.go pkg/headers/doc.go pkg/headers/types.go
	rm -f pkg/headers/headers.go

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:

.PHONY: dep
dep:
	dep ensure -update

.PHONY: install
install: $(SOURCES)
	go install -v  ./cmd/$(NAME)

.PHONY: headers
headers:
	c-for-go -out pkg/ $(HEADERS_CONFIG)
	# Let's remove bad extra fields that c-for-go adds
	sed -i '/ref[a-f0-9].*\|allocs[a-f0-9].*/d' pkg/headers/types.go
	# replace uint by uint64
	sed -i -e "s/uint$$/uint64/g" pkg/headers/types.go
	# this file adds extra methods for our types let's remove them
	rm pkg/headers/cgo_helpers.go

.PHONY: build
build: headers bin/$(EXECUTABLE)

bin/$(EXECUTABLE): $(SOURCES)
	env GOOS=linux GOARCH=amd64 go build -i -v  -o $@ ./cmd/$(NAME)

.PHONY: docs
docs:
	hugo -s docs/