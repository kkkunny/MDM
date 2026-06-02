# AGENTS.md

## Architecture

- Go backend (Echo v5), Flutter web frontend, SQLite database
- Two download engines integrated: Xunlei (XL) and qBittorrent (QB)
- DB models (`dal/db/po/`) and query layer (`dal/db/query/`) are **generated** by `gorm.io/gen` — do not hand-edit
- Protobuf IDL at `model/idl/vo/task.proto` generates Go (`model/vo/`) and Dart (`web/lib/models/vo/`) code

## Commands

```sh
# Install codegen deps (protoc-go-inject-tag)
make install_deps

# Generate Go + Dart from protobuf IDL, then gorm/gen DB models
make gen_idl

# Build Flutter web → static/ (requires flutter SDK)
make gen_web

# Build Go binary (release mode)
CGO_ENABLED=1 go build --tags release -o output/mdm .

# Run locally (dev mode, no build tags — port 8080, DB at ./mdm.db)
go run .
```

## Build tags

- `release` tag controls production behavior:
  - Without `release`: port `8080`, DB at `./mdm.db`, download dir `/mnt/data/downloads`, log level `DEBUG`
  - With `release`: port `80`, DB at `/config/mdm.db`, download dir `/downloads`, log level `INFO`

## Required environment variables

| Variable | Used by |
|---|---|
| `XL_ADDR`, `XL_DID` | Xunlei client (`dal/xl/`) — panics if unset |
| `QB_ADDR`, `QB_USER`, `QB_PASS` | qBittorrent client (`dal/qb/`) — panics if unset |
| `LOG_LEVEL` | Controls log verbosity (DEBUG/TRACE/INFO/WARN/KEYWORD/ERROR/PANIC) |
| `task_download_completed_fallback_addr` | Webhook POSTed on task download completion |

## Key conventions

- Xunlei download tasks name format: `[[category]]|name` — category is parsed out in `model/dto/task.go:60`
- Task IDs are prefixed: `XL|` for Xunlei, `QB|` for qBittorrent (can also be raw DB ID)
- DB connection uses lazy init (`lazy.Getter`) — calling `db.ClientGetter()` or `db.QueryGetter()` triggers DB open + auto-migrate
- Custom Echo JSON serializer (`util.ProtobufJsonEchoSerializer`) handles protobuf messages with `UseProtoNames: true`
- Error middleware (`middleware/response.go`) catches `util.HttpError` and converts to plain text status responses
- All Go errors use `stlerr.ErrorWrap` / `stlerr.ErrorWith` from the `stl` library — never return bare errors
- No test suite exists yet

## Codegen gotchas

- `model/vo/task.pb.go` is generated — if you need a new field, add it to `model/idl/vo/task.proto` and run `make gen_idl`
- `dal/db/po/tasks.gen.go` is generated — table schema changes require running `make gen_idl` (which rebuilds via `cmd/main.go`)
- Run `make gen_idl` before `make gen_web` if proto schemas changed
