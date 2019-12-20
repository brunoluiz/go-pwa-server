DOCKER_REPO:=brunoluiz
PROJECT_NAME:=go-pwa-server

## Prepare version tags
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
VERSION:=$(shell git rev-parse HEAD)
DOCKER_TAG:=$(subst :,-,$(subst /,-,$(GIT_BRANCH)))
ifeq ($(GIT_BRANCH), master)
	DOCKER_TAG := latest
endif

#
# Golang tooling
#
.PHONY: test
test:
	go test -race ./...

build:
	go build -o ./bin/go-pwa-server ./cmd

#
# Docker tooling
#
docker-login:
	docker login --username $(DOCKER_HUB_USER) --password=$(DOCKER_HUB_PASSWORD)

docker-build:
	docker build -t $(DOCKER_REPO)/$(PROJECT_NAME):local .

docker-push:
	docker push $(DOCKER_REPO)/$(PROJECT_NAME)

docker-run:
	docker run -p 80:80 \
		--env-file .env.sample \
		-v $(PWD)/test/static:/static \
		$(DOCKER_REPO)/$(PROJECT_NAME):local

docker-publish: docker-login
	docker build -t $(DOCKER_REPO)/$(PROJECT_NAME):$(VERSION) \
		-t $(DOCKER_REPO)/$(PROJECT_NAME):$(DOCKER_TAG) \
		$(PROJECT_NAME) .
	docker push go-pwa-server
