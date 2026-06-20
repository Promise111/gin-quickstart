# gin-quickstart

A personal sandbox for learning [Gin](https://github.com/gin-gonic/gin) and Go HTTP fundamentals. This is not a production app — it exists to practice routing, form data, multipart uploads, and a standard Go project layout.

## What this covers

- Path params and query strings
- `application/x-www-form-urlencoded` form data (including nested keys)
- Single and multiple file uploads
- Request body size limits with `http.MaxBytesReader`
- Multipart parsing and `MaxMultipartMemory`
- Randomized filenames for uploaded files
- API versioning with Gin route groups (`/api/v1/...`)

## Project structure

```
cmd/api/          Entry point — starts the server
internal/
  router/         Route registration and Gin engine setup
  handler/        HTTP handlers (request in, JSON out)
  utils/          Shared helpers and constants (paths, upload limits)
files/            Single uploads (gitignored)
multiple/         Multi-file uploads (gitignored)
```

## Requirements

- Go 1.25+

## Run locally

```bash
go run ./cmd/api/
```

The server starts on `http://localhost:8080` by default.

## API routing

All endpoints live under a versioned prefix:

```
/api/v1/...
```

Path prefixes are defined in `internal/utils/const.go` (`APIPrefix`, `V1Prefix`) and mounted in `internal/router/router.go` using Gin route groups. When a v2 is needed later, add `registerV2(api.Group("/v2"))` without changing existing v1 handlers.

## API endpoints

Base URL: `http://localhost:8080/api/v1`

### `GET /:name`

Fetch dummy user data by path param.

```bash
curl "http://localhost:8080/api/v1/Promise?gender=m"
```

### `POST /user`

Create a user from form data. Supports query params, post form fields, and nested map keys.

```bash
curl -X POST "http://localhost:8080/api/v1/user?gender=m&ids[course1]=101&ids[course2]=102" \
  -d "name=Promise" \
  -d "age=30" \
  -d "names[course1]=Go Basics" \
  -d "names[course2]=Gin Intro"
```

### `PUT /upload`

Upload a single file (`profilePic`). Max request body: **1 MB**.

```bash
curl -X PUT "http://localhost:8080/api/v1/upload" \
  -F "profilePic=@/path/to/image.jpg"
```

Saved to `./files/` with a randomized filename.

### `PUT /multiple-upload`

Upload multiple files (`files`). Max request body: **4 MB**.

```bash
curl -X PUT "http://localhost:8080/api/v1/multiple-upload" \
  -F "files=@/path/to/image1.jpg" \
  -F "files=@/path/to/image2.png"
```

Saved to `./multiple/` with randomized filenames.

## Upload limits

| Setting | Value | Purpose |
|---------|-------|---------|
| `utils.MaxUploadSizeSingle` | 1 MB | Max body size for `/api/v1/upload` |
| `utils.MaxUploadSizeMultiple` | 4 MB | Max body size for `/api/v1/multiple-upload` |
| `router.MaxMultipartMemory` | 8 MB | How much multipart file data stays in RAM before spilling to disk |

`http.MaxBytesReader` limits the **entire request body** (files + form fields). It is not a per-file-only limit.

## Notes

- Responses use a consistent `{ status, message, ... }` JSON shape with placeholder data.
- There is no database, authentication, or business logic layer.
- Uploaded files are stored on the local filesystem and are excluded from git via `.gitignore`.

## License

Personal learning project — use however you like.
