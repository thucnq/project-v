# Workflow Service

## DB Schemas


## Implement new entity flow

### Pre-require:
- Need install migrate: https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/
- Need install sqlc: https://docs.sqlc.dev/en/latest/overview/install.html

### Steps:
- Define Entity in `./internal/model/entity
- run commend below to gen new migration sql file:
```bash
    make migrate-create name=init_entity
``` 

- Define table, index,.. in sql file above
- Run
```bash
    go run ./cmd/migrate/main.go
```
- Run
```bash
    make gen-sql
```

### Generate api doc:
- swag init -g /cmd/server/main.go
- swag fmt

