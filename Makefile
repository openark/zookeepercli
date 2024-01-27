.PHONY: clean

BINARY_NAME    ?= zookeepercli
VERSION        ?= $(shell git describe --long --tags --always --dirty --abbrev=10)
GOLANG_VERSION ?= 1.21-alpine

CURRENT_DIR := $(shell pwd)

DOCKER ?= docker

.prepare_cmd:
	printf "go build -o ./bin/%s -ldflags \"-X 'main.Version=%s'\"\n" "$(BINARY_NAME)" "$(VERSION)" > .prepared_cmd

# build bin in docker, package in docker and build the docker image
all_in_docker: package_in_docker build_docker_image
	@echo "All done"

# build bin and package locally
all: package
	@echo "All done"

build: .prepare_cmd
	sh .prepared_cmd

package: build
	@ /bin/sh ./scripts/build_package.sh
	@ echo "=== Packaged files: ==="
	@ echo "==== in ./pkgs_out ===="
	@ ls -1 ./pkgs_out
	@ echo "======================="

build_in_docker: .prepare_cmd
	${DOCKER} run --rm -t -v "$(CURRENT_DIR):/app" -w /app \
		-e CGO_ENABLED=0 \
		golang:$(GOLANG_VERSION) /bin/sh .prepared_cmd

package_in_docker: build_in_docker
	${DOCKER} run --rm -t -v "$(CURRENT_DIR):/app" -w /app \
		-e CGO_ENABLED=0 \
		golang:$(GOLANG_VERSION) \
			/bin/sh -c \
				"/bin/sh ./scripts/prepare_container.sh && /bin/sh ./scripts/build_package.sh"
		@ echo "=== Packaged files: ==="
		@ echo "==== in ./pkgs_out ===="
		@ ls -1 ./pkgs_out
		@ echo "======================="

build_docker_image: build_in_docker
	${DOCKER} build -t zookeepercli:$(VERSION) -t zookeepercli:latest .

clean:
	rm -rf ./bin
	rm -rf ./pkgs_out
	rm -rf .prepared_cmd
