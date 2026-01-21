CONTAINER := podman-compose
DEV_SERVICES := dev
PROD_SERVICES := prod

.PHONY: up upp build buildp shell down logs vendor shellp

up:
	$(CONTAINER) up $(DEV_SERVICES)

upp:
	$(CONTAINER) up -d $(PROD_SERVICES)

build:
	$(CONTAINER) build dev

buildp:
	$(CONTAINER) build prod

shell:
	$(CONTAINER) exec dev bash

shellp:
	$(CONTAINER) exec prod bash

down:
	$(CONTAINER) stop

logs:
	$(CONTAINER) logs -f

# Create vendor directory in your project.
# Good for LSP server for your editor.
vendor:
	$(CONTAINER) exec dev go mod vendor

.DEFAULT_GOAL := up
