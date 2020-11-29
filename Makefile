.PHONY: build all
.SILENT: build

ARCH ?= amd64
OPERATOR_HOME ?= /go/src/github.com/jinghzhu/KubernetesCRDOperator
ALL_ARCH = amd64 dawrin
GO_VERSION ?= 1.9
GO_IMAGE ?= golang
GO_BIN ?= $(OPERATOR_HOME)/bin/$(ARCH)
CMDS = github.com/jinghzhu/KubernetesCRDOperator/cmd/operator

all: build

build:
	for command in $(CMDS) ; do \
		echo "building $$command......."; \
		docker run --rm -u $$(id -u):$$(id -g) -v $$(pwd):$(OPERATOR_HOME) \
			-it $(GO_IMAGE):$(GO_VERSION) \
			/bin/sh -c "\
				mkdir -p $(GO_BIN) && \
				GOBIN=$(GO_BIN) go install $$command " && \
		BIN=$$(basename $$command) && echo "Generated bin/$(ARCH)/$$BIN" ; \
	done
