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

# Clean binary and dynamically generated strucutres.
.PHONY: clean-all
clean-all:
	rm -f pkg/headers/*

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

# Generate go data structures from C headers no need to do it for every compil
# just if there are changes on the Natasha headers
# For more informations see natasha_headers.yml
.PHONY: headers
headers:
	c-for-go -out pkg/ $(HEADERS_CONFIG)
	# Let's remove bad extra fields that c-for-go adds
	sed -i '/ref[a-f0-9].*\|allocs[a-f0-9].*/d' pkg/headers/types.go
	# replace uint by uint64
	sed -i -e "s/uint$$/uint64/g" pkg/headers/types.go
	# Remove unused headers
	rm pkg/headers/cgo_helpers.go
	rm pkg/headers/cgo_helpers.h
	# Remove C includes
	sed -i '/^#include.*/d' pkg/headers/*

.PHONY: build
build: bin/$(EXECUTABLE)

bin/$(EXECUTABLE): $(SOURCES)
	go build -i -v  -o $@ ./cmd/$(NAME)

.PHONY: docs
docs:
	hugo -s docs/