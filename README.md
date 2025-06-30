# gRPC Ecommerce Backend System

## Microservice Architecture

### Account Service 
Service: `/account`

Database: PostgreSQL

Tool: `buf`, `direnv`

### Catalog Service
Service: `/catalog`

Database: ElasticSearch

Tool: `buf`, `direnv`

### Order Service
Service: `/order`

Database: PostgreSQL

Tool: `buf`, `direnv`

## GraphQL Gateway
Service: `/graphql`

Tool: `gqlgen`

## Docker Containerization
`docker-compose.yaml`

`/vendor`

# Bazel

### Top-level
`BUILD.bazel`, `MODULE.bazel`

### Account Service
Proto, gRPC: `/account/account/v1/BUILD.bazel`

Server: `/account/cmd/account/BUILD.bazel`

Library: `/account/BUILD.bazel`

### Order Service
Proto, gRPC: `/order/order/v1/BUILD.bazel`

Server: `/order/cmd/order/BUILD.bazel`

Library: `/order/BUILD.bazel`

### Catalog Service
Proto, gRPC: `/catalog/catalog/v1/BUILD.bazel`

Server: `/catalog/cmd/catalog/BUILD.bazel`

Library: `/catalog/BUILD.bazel`

### GraphQL
`/graphql/BUILD.bazel`

## Docker with Bazel

### Top-level
`docker-compose-b.yaml`

### Account Service
`/account/Dockerfile`, `/account/docker_build.sh`

### Order Service
`/order/Dockerfile`, `/order/docker_build.sh`

### Catalog Service
`/catalog/Dockerfile`, `/catalog/docker_build.sh`

### GraphQL
`/graphql/Dockerfile`, `/graphql/docker_build.sh`