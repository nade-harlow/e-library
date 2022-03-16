migrateup:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose up

migratedown:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose down

mock:
	mockgen -source=app/models/db.go -destination=app/mocks/mock_db.go -package=db


run: |
	gofmt -w .
	go run main.go