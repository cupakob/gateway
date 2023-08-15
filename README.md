# BaseMind.AI Monorepo

This is a TypeScript and Golang monorepo, hosting the BaseMind.AI backend services.

## Structure

-   `ts-services` - TypeScript based microservices.
-   `ts-shared` - TypeScript shared code.
-   `go-services` - Golang based microservices.
-   `go-shared` - Golang shared code.
-   `sql` - SQL schema and query files from which we generate the DB DAL (Data Access Layer) and migrations.
-   `docker` - Dockerfiles.
-   `.secrets` - secret values that are gitignored.

The repository root has all tooling and other configurations.

## Installation

1. Install [TaskFile](https://taskfile.dev/) and the following prerequisites:

    - Node >= 20
    - Go >= 1.21
    - Docker >= 24.0
    - Python >= 3.11

    Notes: - Its recommended to use [nvm](https://github.com/nvm-sh/nvm) to manage Node versions. - Its recommended to use [pyenv](https://github.com/pyenv/pyenv) to manage Python versions.

2. Execute the setup task with:

```shell
task deps:setup
```

### TaskFile

We use [TaskFile](https://taskfile.dev/) to orchestrate tooling and commands.
To see all the available commands run:

```shell
task --list
```

### Docker

We use docker for both deployment and local development. There is a `docker-compose.yaml` file in the repository's root
which should be used for local development.

#### Docker Compose Commands

-   `docker compose up --build` to build and start all microservices.
-   `docker compose up --build <service-name>` to build and start a single service.

Note: add `--detach` to run docker compose in the background.

### Development Database

We use Postgres as our database, to bring up the database for local development, execute `task db:up`.

### SQLC

We use SQLC to generate typesafe GO out of SQL queries.

1. install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) on your machine.
2. update the [schema.sql](sql/schema.sql) or [queries.sql](sql/schema.sql) file under the `sql` directory.
3. execute `sqlc generate` to generate the typesafe GO code.

### DB Migrations

We use [atlas](https://github.com/ariga/atlas) for migrating the database.

1. install [atlas](https://github.com/ariga/atlas) on your machine.
2. migrate the database locally for development by executing `task migrations:apply` in the repository root.

Note: you can create new migrations using the task command: `task migrations:create -- <migration_name>`
