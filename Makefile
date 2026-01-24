CONTAINER := podman-compose

# Create vendor directory in your project.
# Good for LSP server for your editor.
vendor:
	$(CONTAINER) exec dev go mod vendor

# Helpful command for development when you need to restart server
restart:
	@$(CONTAINER) restart dev && \
		$(CONTAINER) up dev -d && \
		$(CONTAINER) exec dev npm run dev

.DEFAULT_GOAL := restart
