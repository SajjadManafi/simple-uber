postgres:
	sudo docker run --name postgres-SU -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	sudo docker exec -it postgres-SU createdb --username=root --owner=root simple_uber

dropdb:
	sudo docker exec -it postgres-SU dropdb simple_uber

.PHONY: postgres createdb dropdb