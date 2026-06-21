# Production Projects

Four production-grade projects demonstrating enterprise patterns from beginner to staff level.

## Projects

| # | Project | Key Technologies |
|---|---------|-----------------|
| 1 | [Enterprise REST API](project-1-enterprise-api/) | Auth, PostgreSQL, Redis, Docker, CI/CD, metrics |
| 2 | [Realtime Chat](project-2-realtime-chat/) | WebSocket, Redis Pub/Sub, horizontal scaling |
| 3 | [E-Commerce Backend](project-3-ecommerce/) | Payments, inventory, orders, CQRS |
| 4 | [Microservices Platform](project-4-microservices/) | gRPC, Kafka, distributed tracing, Kubernetes |

## Architecture Overview

```mermaid
flowchart TB
    subgraph Project1[Enterprise REST API]
        API[HTTP API] --> AUTH[Auth Middleware]
        AUTH --> UC[Use Cases]
        UC --> REPO[Repository]
        REPO --> PG[(PostgreSQL)]
        UC --> REDIS[(Redis Cache)]
    end

    subgraph Project2[Realtime Chat]
        WS[WebSocket Server] --> HUB[Connection Hub]
        HUB --> PUBSUB[Redis Pub/Sub]
    end

    subgraph Project4[Microservices]
        GW[API Gateway] --> SVC1[User Service]
        GW --> SVC2[Order Service]
        SVC1 --> KAFKA[Kafka]
        SVC2 --> KAFKA
    end
```

## Getting Started

Each project has its own README with setup instructions:

```bash
cd project-1-enterprise-api
docker compose up -d
go run ./cmd/server/
```

## Production Checklist

Every project includes:

- [ ] Structured logging with correlation IDs
- [ ] Graceful shutdown
- [ ] Health and readiness probes
- [ ] Prometheus metrics
- [ ] OpenTelemetry tracing
- [ ] Docker multi-stage builds
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Integration tests
