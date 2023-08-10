# backend-services

This repository contains the BaseMind.AI backend services. Each service is a Golang microservice.

We use [go-chi](https://github.com/go-chi/chi) as an HTTP 1.1 REST router and [grpc](https://grpc.io/) for
service to service communication.

For databases we use Postgres SQL for which we are writing schemas and queries in SQL. We use [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
to generate a Golang DAL (Data Access Layer) from the SQL, and [atlas](https://github.com/ariga/atlas) to manage
migrations.

## setup

### Docker

We use docker for both deployment and local development. There is a `docker-compose.yaml` file in the repository's root
which should be used for local development.

#### Docker Compose Commands

- `docker compose up --build` to build and start all microservices.
- `docker compose up --build <service-name>` to build and start a single service.

Note: ass `--detach` to run docker compose in the background.

### TaskFile

We use [taskfile](https://taskfile.dev/) to manage local commands.

1. Install [taskfile](https://taskfile.dev/) on your machine.
2. Execute one of the defined commands.

### pre-commit

We use [pre-commit](https://pre-commit.com/) to orchestrate linting and formatting.

1. Install [pre-commit](https://pre-commit.com/) on your machine.
2. install the pre-commit hooks by executing `pre-commit install` in the repository root.

#### pre-commit commands

- `pre-commit autoupdate` to update the hooks.
- `pre-commit run --all-files` to execute against all files in the repository.

Note: do the above commands often.

### Development Database

We use Postgres as our database, to bring up the database for local development, execute `docker compose up db`.

### SQLC

We use SQLC to generate typesafe GO out of SQL queries.

1. install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) on your machine.
2. update the [schema.sql](sql/schema.sql) or [queries.sql](sql/schema.sql) file under the `sql` directory.
3. execute `sqlc generate` to generate the typesafe GO code.

### DB Migrations

We use [atlas](https://github.com/ariga/atlas) for migrating the database.

1. install [atlas](https://github.com/ariga/atlas) on your machine.
2. migrate the database locally for development by executing `task apply-migrations` in the repository root.

Note: you can create new migrations using the task command: `task create-migration -- <migration_name>`
