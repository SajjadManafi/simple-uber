postgres:
	sudo docker run --name postgres-SU -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	sudo docker exec -it postgres-SU createdb --username=root --owner=root simple_uber

dropdb:
	sudo docker exec -it postgres-SU dropdb simple_uber

migrateup:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/simple_uber?sslmode=disable" -verbose up

migrateup1:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/simple_uber?sslmode=disable" -verbose up 1

migratedown:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/simple_uber?sslmode=disable" -verbose down

migratedown1:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/simple_uber?sslmode=disable" -verbose down 1

gotestcover:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 gotestcover