RUNNER = podman-compose

.PHONY: userenv
userenv:
	USER_ID=$(id -u)
	GROUP_ID=$(id -g)

.PHONY: down
down:
	$(RUNNER) down

.PHONY: build
build:
	$(RUNNER) build

.PHONY: db
db:
	$(RUNNER) exec db bash

.PHONY: stop
stop:
	$(RUNNER) stop

.PHONY: shell
shell:
	$(RUNNER) exec app bash

.PHONY: up
up:
	$(MAKE) userenv
	@if [ ! -d "node_modules" ]; then $(RUNNER) run --rm  app npm install; fi
	$(RUNNER) up -d

.PHONY: dev
dev: up
	$(RUNNER) exec app bash -c "npm run dev"

.PHONY: restart
restart: down dev

.DEFAULT_GOAL := shell
