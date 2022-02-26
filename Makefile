migrateup:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose up

migratedown:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose down

run:
	go run main.go