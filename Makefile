postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:14-alpine

createdb: 
	docker exec -it postgres createdb --username=root --owner=root bank   

dropdb: 
	docker exec -it postgres dropdb bank

migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5433/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5433/bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown