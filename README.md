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
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations create migration_file_name

# apply all available migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations up

# roll back all migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations down

# roll back the most recently applied migration, then run it again.
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations redo

# run down and then up command
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations reset

# show the current migration version
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations version

# apply the next n migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate +1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate +2
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate +n

# roll back the previous n migrations
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate -1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate -2
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations migrate -n

# go to specific migration
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations goto 1
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations goto 10
migrate -url cassandra://127.0.0.1:9042/task_manager -path ./backend/migrations goto v
```

# React app
## Developing

run
```sh
$ npm install
```

start webpack dev server
```sh
$ npm start
```

serve static html
```sh
$ go run dev.go
```
go to localhost:8080/static/


## Get prebuild app

run
```sh
$ npm run build:prod
```
in your index.html change
```html
<script src="http://localhost:3000/bundle.js"></script>
```
to
```sh
<script src="bundle.js"></script>
```
