.PHONY: down
down:
	@podman-compose down

.PHONY: build
build:
	@podman-compose build

.PHONY: db
db:
	@podman-compose exec db bash

.PHONY: backup
backup:
	podman-compose exec -T db mysql -u serhii -ppassword serhii < backup.sql

.PHONY: stop
stop:
	@podman-compose stop

.PHONY: shell
shell:
	@podman-compose exec app bash

.PHONY: run
run:
	@podman-compose exec app go run ./cmd/blog

.PHONY: up
up:
	@USER_ID=$(id -u)
	@GROUP_ID=$(id -g)
	@if [ ! -d "node_modules" ]; then podman-compose run --rm  app npm install; fi
	@podman-compose up -d
	@echo "ENTERING CONTAINER!"
	@$(MAKE) shell

.PHONY: dev
dev:
	@podman-compose exec app npm run dev

.PHONY: restart
restart: down dev

.DEFAULT_GOAL := shell
