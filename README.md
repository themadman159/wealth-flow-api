# Golang API Project

### Description

This is a Golang API project template with Docker support and environment configuration.

### Prerequisites

- Go (latest version recommended)
- Docker and Docker Compose
- Git

## Getting Started

1. Clone the repository
```
git clone https://github.com/themadman159/go-api.git
```

2. Install dependencies
```
go mod tidy -e
```

3. Set up environment configuration
```
cp .env-example .env
```
```
cp docker-compose-example.yaml docker-compose.yaml
```

4. Start Docker services
```
docker compose up -d
```

5. Run the Server
```
go run cmd/main.go
```






