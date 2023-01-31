.PHONY: postgres admniner migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=postgres -d postgres
adminer:
	docker run --rm -ti --network host -d adminer
migrate: 
	migrate -source file://migrations \
	-database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up
