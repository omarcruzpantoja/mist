

up: down
	docker compose build
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

# API
logs:
	docker logs -f -n 100 ${SERVICE_NAME}-api

console:
	docker exec -it ${SERVICE_NAME}-api /bin/bash

# DB
initialize-db:
	docker exec -it ${SERVICE_NAME}-scylla cqlsh -u 'cassandra' -p 'cassandra' -f src/scylladb/scripts/initialize.cql

db-logs:
	docker logs -f -n 100 ${SERVICE_NAME}-scylla

db-console:
	docker exec -it ${SERVICE_NAME}-scylla cqlsh

# DB Migrations
generate-migration gm:
	docker exec -it ${SERVICE_NAME}-api migrate create -ext sql -dir src/scylladb/migrations -seq $$message

migrate:
	docker exec -it ${SERVICE_NAME}-api migrate -path src/scylladb/migrations -database "cassandra://cassandra:cassandra@user-service-scylla:${SCYLLA_PORT}/$(SCYLLA_KEYSPACE)" up

downgrade:
	docker exec -it ${SERVICE_NAME}-api migrate -path src/scylladb/migrations -database "cassandra://cassandra:cassandra@user-service-scylla:${SCYLLA_PORT}/$(SCYLLA_KEYSPACE)" down ${count}

db-version:
	docker exec -it ${SERVICE_NAME}-api migrate -path src/scylladb/migrations -database "cassandra://cassandra:cassandra@user-service-scylla:${SCYLLA_PORT}/$(SCYLLA_KEYSPACE)" version


db-force:
	docker exec -it ${SERVICE_NAME}-api migrate -path src/scylladb/migrations -database "cassandra://cassandra:cassandra@user-service-scylla:${SCYLLA_PORT}/$(SCYLLA_KEYSPACE)" force ${version}

# LINT
lint: 
	docker exec -it ${SERVICE_NAME}-api golangci-lint run