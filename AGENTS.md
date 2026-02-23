# Serhii Cho Blog - Agent Guidelines

## Project Overview
Personal blog rewritten from Laravel to Go using Textwire templating. Learning project exploring Go with a real-world application using Textwire (custom templating language by the same author).

Live: https://serhiicho.com

## Stack
- **Go 1.26** with standard library http
- **Textwire v3** for templating (github.com/textwire/textwire/v3)
- **MySQL 8.0** with go-sql-driver/mysql
- **Vue 3 + TypeScript + Tailwind CSS v4**
- **Vite** build tool
- **Podman** for containers

## Project Structure
```
cmd/blog/main.go           # Entry point
internal/
  config/                  # Config types (JSON-based)
  database/
    database.go            # Repository interfaces
    mysql/                 # MySQL implementations
  http/
    server.go              # HTTP server & routing
    handlers.go            # Route handlers
    textwire.go            # Template engine config
  model/                   # Domain models with db/json tags
templates/
  layouts/                 # Base layouts (use ~ prefix)
  views/                   # Page views
  components/              # Reusable components
    home/                  # Homepage components
    icons/                 # SVG icons
    partials/              # Shared partials
    posts/                 # Post components
web/
  css/                     # Stylesheets (Tailwind)
  ts/                      # TypeScript (Vue apps)
  articles/                # Markdown article files
  storage/                 # Static assets
public/                    # Public files
config/config.json         # App configuration
```

## Commands

### Development (via Makefile)
```bash
make up              # Start containers and enter shell (default)
make shell           # Enter running app container
make down            # Stop containers
make build           # Build containers
make restart         # Restart with dev mode
make dev             # Run Vite dev server (inside container)
make run             # Run Go app locally (run when you are inside of a container)
make db              # Enter MySQL shell
make backup          # Restore from backup.sql
```

### Go Commands inside app container
```bash
go test ./...                    # Run all tests
go test -run TestName ./path     # Run specific test
go fmt ./...                     # Format code
go vet ./...                     # Static analysis
```

### Frontend inside app container
```bash
npm run dev              # Vite dev server (port 5173)
npm run build            # Production build
```

## Code Conventions

### Go

**Imports** (3 groups: stdlib, external, internal):
```go
import (
    "context"
    "database/sql"
    
    "github.com/joho/godotenv"
    
    "serhii/internal/config"
)
```

**Naming**:
- Exported: PascalCase (`Post`, `UserRepository`)
- Private: camelCase (`db`, `latest()`)
- Files: lowercase (`post.go`, `database.go`)
- Interfaces: suffixed with type (`PostRepository`)

**Struct Tags** (always include both):
```go
type Model struct {
    ID        int       `json:"id" db:"id"`
    Title     string    `json:"title" db:"title"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}
```

**Error Handling**:
- Repositories: return errors to caller
- Handlers: log with `log.Println(err)` and continue (don't crash server)
- Fatal only in `main.go` during initialization

**Pattern**:
- Repository interface defined in `internal/database/database.go`
- Implementation in `internal/database/mysql/{entity}.go`
- Database injected via constructor

### Textwire Templates (.tw files)

**Key Directives**:
- `@use("~main")` - extend layout (~ = templates root)
- `@insert("title", "content")` - define section content
- `@component("~path/component", { props })` - render component
- `@slot` / `@end` - component content block
- `{{ variable }}` - output (auto-escaped)
- `{{ defined(var) ? var : 'default' }}` - conditional with default

**Naming Conventions**:
- Layouts: `~main`, `~admin`
- Components: `~home/section`, `~icons/arrow-right`
- Partials: `~partials/head`, `~partials/navbar`
- Views: no prefix (resolved relative to templates/)

**Global Data** (available in all templates):
Access via `global.` prefix:
- `global.env` - APP_ENV value
- `global.appName` - APP_NAME value
- `global.darkThemeCookie` - theme preference
- `global.manifest.js/css` - Vite manifest entries

### Frontend

**Entry Points**:
- Main: `web/ts/main/app.ts` + `web/css/main/app.css`
- Admin: `web/ts/admin/app.ts` + `web/css/admin/app.css`

**Aliases** (Vite):
- `@/` → `web/ts/main/`
- `@admin/` → `web/ts/admin/`
- `@user/` → `web/ts/user/`
- `@shared/` → `web/ts/shared/`

**Conventions**:
- Vue 3 Composition API with `<script setup>`
- Tailwind v4: use `@import "tailwindcss"` in CSS
- TypeScript strict mode

## Environment
Required variables (see `.env.example`):
- `APP_NAME`, `APP_ENV`, `APP_PORT`
- `DB_HOST`, `DB_PORT`, `DB_DATABASE`, `DB_USERNAME`, `DB_PASSWORD`

## Development Notes
- Textwire auto-reloads templates in development (`FileWatcher: true`)
- MySQL connection pooling configured (max 20 connections)
- Vite HMR on port 5173, app on port 8080
- Public files served from `./public` directory
- Template components receive props as object parameter

## Testing from the container
No tests exist yet. Use standard Go testing:
- Test files: `*_test.go` alongside source files
- Run: `go test ./...`
- Single test: `go test -run TestName ./package`

## Common Patterns

**Adding a New Page**:
1. Create handler in `internal/http/handlers.go`
2. Register route in `internal/http/server.go`
3. Create template in `templates/views/`
4. Use `@use('~main')` to extend `layouts/main.tw` (`~` an alias for `layouts/`)
5. Use `@insert('title', 'Some title')` adn `@insert('content')` to insert things into a layout file

**Adding a Repository Method**:
1. Add interface method to `internal/database/database.go`
2. Implement in `internal/database/mysql/{entity}.go`
3. Return `([]*model.Entity, error)` or `(*model.Entity, error)`

**Adding a Component**:
1. Create `.tw` file in `templates/components/{category}/`
2. Use `@slot` for content injection
3. Access props with `{{ defined(propName) ? propName : 'default' }}`
4. Reference with `~` prefix: `@component("~category/name", { prop: value })` (`~` an alias for `components/`)

## Textwire
- All Textwire files end with `.tw` file extension
- Templates live in `./templates/`
- Components live in `./templates/components/`
- Layouts live in `./templates/layouts/`
- Views live in `./templates/views/`
