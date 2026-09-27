# City Counter

City Counter counts cities whose names start with a requested letter. The project includes a small web interface and a Go HTTP API backed by the GeoNames public search API.

## Requirements

- Docker and Docker Compose for the complete application
- Go 1.27.1 or later for running the backend locally
- Internet access, because the backend queries GeoNames

## Launch with Docker Compose

From the repository root, build and start both services:

```bash
docker compose up --build
```

Open the application at [http://localhost:3000](http://localhost:3000).

The services are exposed at:

- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`
- Backend health check: `http://localhost:8080/health`

Stop the services with:

```bash
docker compose down
```

The backend container has a health check. The frontend waits for the backend to become healthy before starting.

## Backend API

### `GET /health`

Returns `200 OK` with the plain-text body `OK` when the backend is running.

### `GET /api/cities/count`

Counts cities whose names start with the supplied value.

Query parameters:

| Parameter  | Required | Values                     | Default      | Description                                                   |
| ---------- | -------- | -------------------------- | ------------ | ------------------------------------------------------------- |
| `letter`   | Yes      | A letter or prefix         | None         | Value used to match city names. Matching is case-insensitive. |
| `strategy` | No       | `startswith`, `bruteforce` | `startswith` | Counting strategy to use.                                     |

Example:

```bash
curl 'http://localhost:8080/api/cities/count?letter=A'
```

Response:

```json
{
	"letter": "A",
	"count": 123
}
```

The API returns `500 Internal Server Error` with the following JSON shape if GeoNames or the selected strategy fails:

```json
{
	"error": "Server error"
}
```

CORS is enabled for all origins. The frontend uses the same API paths through the nginx reverse proxy.

## Counting strategies

Both strategies cache results in memory for the lifetime of the backend process. A cache is maintained separately for each strategy.

### `startswith` (default)

- Sends a GeoNames query with `name_startsWith` set to the requested value.
- Requests only one result row and uses GeoNames' `totalResultsCount`.
- Usually the most efficient option because the remote service performs the prefix filtering.

### `bruteforce`

- Requests up to 1,000 cities from GeoNames.
- Checks each returned city locally with a case-insensitive prefix comparison.
- Useful for demonstrating local counting, but it may not represent the full GeoNames result set when more than 1,000 cities match.

## Run the backend locally

Start the backend directly from its directory:

```bash
cd backend
go run ./cmd/server
```

The server listens on `http://localhost:8080` by default. Set `PORT` to use another port:

```bash
PORT=9090 go run ./cmd/server
```

The GeoNames URL and username are currently configured in `backend/internal/config/config.go`. The backend needs outbound network access to query that service.

## Run tests

```bash
cd backend
go test ./...
```

## Project structure

- `backend/cmd/server`: HTTP server entry point
- `backend/internal/city`: city API handler, service, repository, cache, and strategies
- `backend/internal/geonames`: GeoNames HTTP client
- `frontend/index.html`: browser UI
- `frontend/nginx.conf`: static file server and backend reverse proxy
- `docker-compose.yml`: local multi-service setup
