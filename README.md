# Go Todo API 

A robust RESTful backend service built with Go, featuring MySQL persistence and Redis caching.

## Tech Stack
*   **Core:** Go 1.21, Gin Web Framework
*   **Data:** MySQL 8, GORM (ORM)
*   **Cache:** Redis (Go-Redis)
*   **Infrastructure:** Docker & Docker Compose

## How to Start
Run the entire stack (App + DB + Redis) with one command:

```bash
make build
```

## API Endpoints

* Create a TODO card
* Get all cards
* Update a card

## Usage Examples

* Create a Card 

```bash
curl -X POST http://localhost:8080/api/cards \
-H "Content-Type: application/json" \
-d '{"name": "Fix Bug", "priority": "High", "due_date": "2025-12-31"}'
```

* Get All

```bash
curl http://localhost:8080/api/cards
```

* Update a card
```bash
curl -X PATCH http://localhost:8080/api/cards/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```
