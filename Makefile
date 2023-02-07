SHELL = /bin/bash
NAME ?= what-is-my-ip

PKGLIST=$(shell go list ./...)
BUILD_GIT_TAG=$(shell git describe --tags --abbrev=0 || echo "TagNotFound")
BUILD_GIT_COMMIT_ID=$(shell git rev-parse HEAD || echo "GitNotFound")
BUILD_REPO_BRANCH=$(shell git symbolic-ref --short HEAD || echo "GitNotFound")
LDFLAGS="\
-s -w \
-X github.com/Dup4/what-is-my-ip/version.BUILD_GIT_TAG=$(BUILD_GIT_TAG) \
-X github.com/Dup4/what-is-my-ip/version.BUILD_GIT_COMMIT_ID=$(BUILD_GIT_COMMIT_ID) \
-X github.com/Dup4/what-is-my-ip/version.BUILD_REPO_BRANCH=$(BUILD_REPO_BRANCH) \
"

all: build

build:
	go build -ldflags $(LDFLAGS) -o output/what-is-my-ip

install:
	go install -ldflags $(LDFLAGS)

test:
	go vet $(PKGLIST)
	go test $(PKGLIST) -race -coverprofile=./unittest-coverage.out

ut:
	go test $(PKGLIST) -race -coverprofile=./unittest-coverage.out

bench:
	go test $(PKGLIST) -run=NOTEST -benchmem -bench=. -cpu=1,2,4,8

clean:
	rm -rf ./output
