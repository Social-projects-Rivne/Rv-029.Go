# Install
1. Install base set to $GOPATH/src/github.com/Social-projects-Rivne/Rv-029.Go

2. Create environment file ".env"
```sh
cp .env.example .env
```
3. Set up correct values in .env

4. Start containers
```sh
docker-compose up -d cassandra react
```

# Migrations

## Create keyspace

```cqlsh
CREATE KEYSPACE task_manager 
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'} 
    AND durable_writes = true;
```

## Usage

```bash
# install
go get github.com/gemnasium/migrate

# create new migration file in path
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations create migration_file_name

# apply all available migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations up

# roll back all migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations down

# roll back the most recently applied migration, then run it again.
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations redo

# run down and then up command
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations reset

# show the current migration version
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations version

# apply the next n migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate +1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate +2
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate +n

# roll back the previous n migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate -1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate -2
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations migrate -n

# go to specific migration
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations goto 1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations goto 10
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./migrations goto v
```