postgres:
	sudo docker run --name postgres-SU -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

.PHONY: postgres