CONTAINER := podman-compose
DEV_SERVICES := dev
PROD_SERVICES := prod

.PHONY: up upp build buildp shell down stop logs vendor shellp ps restart

up:
	$(CONTAINER) up -d $(DEV_SERVICES)

restart:
	$(CONTAINER) restart $(DEV_SERVICES)

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

stop:
	$(CONTAINER) stop

down:
	$(CONTAINER) down

logs:
	$(CONTAINER) logs -f

ps:
	$(CONTAINER) ps

# Create vendor directory in your project.
# Good for LSP server for your editor.
vendor:
	$(CONTAINER) exec dev go mod vendor

.DEFAULT_GOAL := up
