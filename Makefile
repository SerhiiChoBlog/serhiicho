CONTAINER := podman-compose

# Create vendor directory in your project.
# Good for LSP server for your editor.
.PHONY: vendor
vendor:
	$(CONTAINER) exec dev go mod vendor

# Helpful command for development when you need to restart server
.PHONY: restart
restart:
	@$(CONTAINER) restart dev && \
		$(CONTAINER) up dev -d && \
		$(CONTAINER) exec dev npm run dev

.PHONY: vite
vite:
	$(CONTAINER) exec dev npm run dev

.PHONY: go
go:
	$(CONTAINER) up dev -d

.DEFAULT_GOAL := restart
