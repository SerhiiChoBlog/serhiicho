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

.PHONY: stop
stop:
	$(RUNNER) stop

.PHONY: shell
shell:
	$(RUNNER) exec app bash

.PHONY: dev
dev:
	$(MAKE) userenv
	@if [ ! -d "node_modules" ]; then $(RUNNER) run --rm  --service-ports app npm install; fi
	$(RUNNER) run --rm --service-ports app bash -c "npm run dev & go run ./cmd/blog"

.DEFAULT_GOAL := shell
