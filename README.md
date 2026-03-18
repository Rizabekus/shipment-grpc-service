# Shipment gRPC Service

## 1. Running the Service

```bash
git clone git@github.com:Rizabekus/shipment-grpc-service.git
cd shipment-grpc-service
go run ./cmd/server
```
## 2. Running the Tests

```bash
go test -v ./...
```
## 3. Architecture Overview

The project follows principles of Clean Architecture

internal/domain:
describes business rules and domain models. This layer should be independent from everything

internal/application:
Coordinates interactions between domain and infrastructure layers

internal/transport/grpc: 
manages clients requests, takes input, calls functions from application layer and sends back responses to clients 

Request flow:
gRPC (transport) → application (use cases) → domain / infrastructure

## 4. Design Decisions

For the beginning I created in-memory realization of storage, in the future it can be easily extended with other storages like postgres

Clean Architecture was chosen to ensure separation of concerns and testability

## 5. Assumptions

I assumed that the shipment statuses should include: "PENDING" -> "PICKED_UP" -> "IN_TRANSIT" -> "DELIVERED" -> "COMPLETED"

Also I assumed that the statuses should be updated according to specific order without jumping over more than 1 step ahead