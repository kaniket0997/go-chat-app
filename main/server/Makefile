postgresinit:
	docker pull postgres:15-alpine
	docker run --name postgres15 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
	docker exec -it postgres15 createdb --username=postgres --owner=postgres go-chat
postgres:
	docker exec -it postgres15 psql
dropdb:
	docker exec -it postgres15 dropdb go-chat
migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/go-chat?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres migrateup migratedown

