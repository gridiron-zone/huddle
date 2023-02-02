huddle-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag huddlelabs/huddle-env huddle-env

huddle-python:
	docker build --tag huddlelabs/huddle-python huddle-python

huddle-cosmovisor:
	docker build \
		--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
		--build-arg HUDDLE_VERSION=$(HUDDLE_VERSION) \
		--tag huddlelabs/huddle-cosmovisor:$(HUDDLE_VERSION) \
		huddle-cosmovisor

huddle-builder:
	docker build --tag huddlelabs/builder -f huddle-builder/Dockerfile $(CONTEXT)

all: huddle-env huddle-python huddle-cosmovisor huddle-builder

.PHONY: all huddle-env huddle-python huddle-cosmovisor huddle-builder
