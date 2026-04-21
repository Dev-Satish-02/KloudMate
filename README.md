# KloudMate

KloudMate is a minimal real-time log ingestion and processing stack built with Go, Kafka, Redis, ClickHouse, and Grafana.

![KloudMate dashboard](./Dashboard.png)

## What it does

- **Producer**: generates structured JSON logs and publishes them to Kafka topic `logs`
- **Backend**: consumes from Kafka, stores logs in ClickHouse, maintains rolling per-service error counters in Redis, and prints a basic alert when the error rate spikes
- **Grafana**: runs alongside the stack and can be configured to query ClickHouse for dashboards

## Architecture

1. `producer` → Kafka (`logs`)
2. `backend` consumer (`group-1`) reads from Kafka
3. `backend` writes to:
   - **ClickHouse** table `logs` (durable storage / analytics)
   - **Redis** keys `error_count:<service>` (real-time counters, TTL 60s)

## Repo layout

```
.
├── backend/               # Go consumer + HTTP API
├── producer/              # Go log producer
├── dashboards/            # (optional) dashboard assets
├── docker-compose.yml     # local infrastructure: Kafka/Redis/ClickHouse/Grafana
└── clickhouse-users.xml   # ClickHouse user config (default user, no password)
```

## Prerequisites

- Docker + Docker Compose
- Go (see `backend/go.mod` and `producer/go.mod`)

## Services and ports

- **Kafka**: `localhost:9092`
- **Redis**: `localhost:6379`
- **ClickHouse (HTTP)**: `localhost:8123`
- **ClickHouse (native)**: `localhost:9000`
- **Grafana**: `localhost:3000`
- **Backend API**: `localhost:8080`

## Quickstart (local)

Start the infrastructure:

```bash
docker compose up -d
```

Create the ClickHouse `logs` table (inside the ClickHouse container):

```bash
docker exec -it clickhouse clickhouse-client --multiquery --query "
CREATE TABLE IF NOT EXISTS logs
(
  timestamp DateTime,
  service   String,
  level     String,
  latency   Float32
)
ENGINE = MergeTree
ORDER BY timestamp;
"
```

Run the backend:

```bash
cd backend
go run .
```

Run the producer (in another terminal):

```bash
cd producer
go run .
```

Validate:

```bash
curl -sS localhost:8080/health
```

## API

Currently implemented:

- `GET /health`: returns `OK`

## Grafana

Open Grafana at `http://localhost:3000`.

- The Compose file installs the ClickHouse datasource plugin (`grafana-clickhouse-datasource`).
- Add a ClickHouse datasource pointing to `http://clickhouse:8123`.

## Notes

- The Go services currently connect to Kafka/Redis/ClickHouse via `localhost`, so run `backend` and `producer` on the host while dependencies run in Docker.
- Alerts are currently **console output** based on Redis error counters (threshold: >20 errors within the 60s TTL window).

