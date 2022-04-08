# E-Library
This is a book library application where students check in and then given access to borrow books and return when done.

## Getting Started

1. [Install Go](https://golang.org/doc/install)
2. [Install PostgreSQL](https://www.postgresql.org/download/)
3. Create a database named `e_library`.
4. Install all dependencies by running:


```
go mod tidy
```

5. Run the database migration script in `app/db/migration` by running:

```
make migrateup
```

6. Run the application using:

```
go run main.go
```
## Setting Environmental Variables
An environment variable is a text file containing ``KEY=value`` pairs of your secret keys and other private information. For security purposes, it is ignored using ``.gitignore`` and not committed with the rest of your codebase.

To create, ensure you are in the root directory of the project then on your terminal type:
```
touch .env
```
All the variables used within the project can now be added within the ``.env`` file in the following format:
```
PORT=8080
DB_HOST=127.0.0.1
DB_PORT=<your db port>
DB_USER=<your db username>
DB_PASS=<your db password>
DB_NAME=e_library
```

> Note: This project is built using PostgreSQL so configure the `.env` file using PostgreSQL credentials.

> Click [here](https://harlow-elibrary.herokuapp.com/) see a live demo
