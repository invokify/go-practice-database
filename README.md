# go-practice-database

A clean database client implementation:
- implementing migrations
- generating CRUD operations code using SQLC
- handling transactions
- preventing deadlocks

Reference: [Tech School](https://www.youtube.com/@TECHSCHOOLGURU)

## Get Started

1. Launch Postgres database using Docker
```
make postgres
```

2. Create the database
```
make createdb
```

3. Apply schema migration
```
make migrateup
```

## Isolation Levels

### MySQL

|  | READ UNCOMMITED | READ COMMITED | REPEATABLE READ | SERIALIZABLE |
|----------|----------|----------|----------|----------|
| DIRTY READ    | ✔     | X     | X     | X     |
| NON-REPEATABLE READ    | ✔     | ✔     | X     | X     |
| PHANTOM READ    | ✔     | ✔     | X     | X     |
| SERIALIZATION ANOMALY    | ✔     | ✔     |   ✔   | X     |

- 4 isolation levels
- Locking Mechanism
- Repeatable read

### Postgres

|  | READ UNCOMMITED | READ COMMITED | REPEATABLE READ | SERIALIZABLE |
|----------|----------|----------|----------|----------|
| DIRTY READ    | X     | X     | X     | X     |
| NON-REPEATABLE READ    | ✔     | ✔     | X     | X     |
| PHANTOM READ    | ✔     | ✔     | X     | X     |
| SERIALIZATION ANOMALY    | ✔     | ✔     |   ✔   | X     |

- 3 isolation levels
- Dependencies detection
- Read commited

> 💡 **To keep in mind**: There might be errors, timeout or deadlock > RETRY MECHANISM
