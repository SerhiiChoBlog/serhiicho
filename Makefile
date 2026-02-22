RUNNER = podman-compose

.PHONY: down
down:
	@$(RUNNER) down

.PHONY: build
build:
	@$(RUNNER) build

.PHONY: db
db:
	@$(RUNNER) exec db bash

.PHONY: stop
stop:
	@$(RUNNER) stop

.PHONY: shell
shell:
	@$(RUNNER) exec app bash

.PHONY: run
run:
	@go run ./cmd/blog

.PHONY: up
up:
	@USER_ID=$(id -u)
	@GROUP_ID=$(id -g)
	@if [ ! -d "node_modules" ]; then $(RUNNER) run --rm  app npm install; fi
	@$(RUNNER) up -d
	@echo "ENTERING CONTAINER!"
	@$(MAKE) shell

.PHONY: dev
dev:
	@$(RUNNER) exec app npm run dev

.PHONY: restart
restart: down dev

.DEFAULT_GOAL := shell
