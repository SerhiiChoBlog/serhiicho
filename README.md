# Serhii Go Blog

## Development with Podman
Build a project for development using Dockerfile.dev file:
```bash
make build
```

Start a development server with autorefresh:
```bash
make up
```

Creates a `vendor` directory in your root project. Good for LSP:
```bash
make vendor
```

Enter a development container:
```bash
make shell
```

## Production
Build image for production using Dockerfile.prod file:
```bash
make buildprod
```

Start a production server on port `80`:
```bash
make upprod
```

Enter a production container:
```bash
make shellprod
```
