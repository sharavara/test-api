PHONY: build latest rmi build-dev build-staging build-prod

IMAGE ?= vincentvega/test-api
VERSION = $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")

build-dev build:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE):dev .

build-staging:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE):master .

build-prod:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE):latest .

rmi:
	docker rmi $(IMAGE):$(VERSION)

