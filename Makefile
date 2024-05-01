USER = root
DB_PASSWORD = secret
DB_NAME = simple_bank
DB_PORT = 5432
DB_URL = localhost
DOCKER_NAME = simplebank-postgres-1

createdb:
	docker exec -it $(DOCKER_NAME) createdb --username=$(USER) --owner=root $(DB_NAME)

dropdb:
	docker exec -it $(DOCKER_NAME) dropdb $(DB_NAME)

migration_up:
	goose -dir db/migration postgres "postgresql://$(USER):$(DB_PASSWORD)@$(DB_URL):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migration_down:
	goose -dir db/migration postgres "postgresql://$(USER):$(DB_PASSWORD)@$(DB_URL):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

conn_db:
	psql -U $(USER) -h $(DB_URL) -p $(DB_PORT) -d $(DB_NAME)
	
.PHONY: createdb dropdb migration_up migration_down sqlc test conn_db

# -verbose up