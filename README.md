# Chirpy

This is the project of the Boot.dev course [Learn HTTP Servers in Go](https://www.boot.dev/courses/learn-http-servers-golang)

It's an API inspired in Twitter (in this case Chirpy)

## Install and Run

So, I guess you will need to

1. Install postgres

2. Create a database and name it chirpy

3. Install goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

4. Clone the project

```bash
git clone git@github.com:kasteion/chirpy.git
```

5. Copy the `.env.example` file into a `.env` file and update the `.env`

```bash
cp .env.example .env
```

6. Copy `sql/schema/.env.example` into `sql/schema/.env` and update it with your postgres connection string

```bash
cp sql/schema/.env.example sql/schema/.env
```

7. `cd` into `sql/schema` and run migrations with goose

```bash
cd sql/schema
goose postgres up
```

8. Then you can execute the project (from the root)

```bash
go run .
```

9. Queries are generated with [sqlc](https://github.com/sqlc-dev/sqlc)
