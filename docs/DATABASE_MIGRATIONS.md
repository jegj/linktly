# Database Migration

## Pre Requesites

### Migrate

Install [migrate](https://github.com/golang-migrate/migrate) with the following instructions

```sh
$curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$apt-get update
$apt-get install -y migrate
```

### Create Database

```sh
psql -h 127.0.0.1 -p 5433 -U linktly_admin -w -c "create database linktlydb;"
```

## Create Migrations

```sh
migrate create -ext sql -dir internal/database/migrations -seq create_linktly_schema
```

## Run Migration

```sh
export POSTGRESQL_URL='postgres://linktly_admin:dfNjlX@localhost:5433/linktlydb?sslmode=disable'
```

### up

```sh
migrate -database ${POSTGRESQL_URL} -path internal/database/migrations up
```

### down

```sh
migrate -database ${POSTGRESQL_URL} -path internal/database/migrations down
```

### force

When migration failed for syntax error or any other error. The migration
becomes dirty and need to run again once the error is fixed

```sh
migrate -database ${POSTGRESQL_URL} -path internal/database/migrations force 1
```
