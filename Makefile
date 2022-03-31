migrateup:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose up

migratedown:
	migrate -path app/db/migration -database "postgresql://postgres:postgres@localhost:5432/e_library?sslmode=disable" -verbose down

mock:
	mockgen -source=app/models/db.go -destination=app/mocks/mock_db.go -package=db

herodeploy:
	psql --host=ec2-54-195-76-73.eu-west-1.compute.amazonaws.com --port=5432 --username=evrhmwwrktikrl --password --dbname=dfemrlacff0t1

heroku pg:
	psql --app harlow-elibrary < app/db/migration/000001_create_app_table.up.sql

run: |
	gofmt -w .
	go run main.go