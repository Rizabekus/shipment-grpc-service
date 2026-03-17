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

internal/domain: holds the main business logic and domain models. This layer should be independent from everything
internal/application: manages all the business logic, is being called from transport layer, calls infrastructure and domain layers
internal/transport/grpc: manages clients requests, takes input, calls functions from application layer and sends back responses to clients 
