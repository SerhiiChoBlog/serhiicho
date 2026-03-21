.PHONY: run
run:
	@go run ./cmd/blog

.PHONY: dev
dev:
	npm run dev

.PHONY: dbup
dbup:
	podman-compose up db -d

.PHONY: dbstop
dbstop:
	podman-compose stop db

.DEFAULT_GOAL := run
