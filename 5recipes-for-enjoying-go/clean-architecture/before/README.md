# before

## Starting DB and Server

```shell
docker-compose up
```

## Migration

[golang-migrate](https://github.com/golang-migrate/migrate)

### installation

See installation page https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

or just run the following command

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### cli usage

```shell
set -x POSTGRESQL_URL 'postgres://wwgt-diary:wwgt-diary@0.0.0.0:5435/wwgt-diary?sslmode=disable'

# Run migrations
migrate -database $POSTGRESQL_URL -path migrations up

# Create new migration files
migrate create -ext sql -dir migrations -seq MIGRATION_FILE_NAME

# Run down migrations
migrate -database $POSTGRESQL_URL -path migrations down
```
