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

# Install Dependencies
## Install glide
https://github.com/Masterminds/glide#install

## Install dependencies
```sh
    glide install
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

# Seeds
Run seeds
```sh
go run backend/main.go db:seed
```

Default admin user email: `user@gmail.com`

Default admin user password: `qwerty1234`

# Golang web server

## Run server
Run server with command. By default port uses :8080

```sh
go run backend/main.go
```


## API Routes List

| Action | Route | Input | Method | Auth |
| ------ | ------ | ------ | ------ | ------ |
| Login | auth/login | JSON{ email, password } | POST | NO |
| Registration | auth/register | JSON{ email, name, surname, password } | POST | NO |
| Forget Password | auth/forget-password | JSON{ email } | POST | NO |
| New Password | auth/new-password | JSON{ token, email, password } | POST | NO |
| Create Project | project/create | JSON{ name } | POST | YES |
| Update Project | project/update/:id | JSON{ name } | PUT | YES |
| Delete Project | project/delete/:id | - | DELETE | YES |
| Projects list | project/list | - | GET | YES |
| Show Project | project/show/:id | - | GET | YES |
| Create Board | project/:project_id/board/create | JSON{ name, desc } | POST | YES |
| Update Board | project/board/update/:board_id | JSON{ name, desc } | PUT | YES |
| Delete Board | project/board/delete/:board_id | - | Delete | YES |
| Boards list | project/:project_id/board/list | - | GET | YES |
| Show Board | project/board/show/:board_id | - | GET | YES |
| Create Sprint | project/board/:board_id/sprint/create | JSON{ goal, name, desc } | POST | YES |
| Update Sprint | project/board/sprint/update/:sprint_id | JSON{ goal, name, desc, status, ...fields } | PUT | YES |
| Delete Sprint | project/board/sprint/delete/:sprint_id | - | Delete | YES |
| Spints list | project/board/:board_id/sprint/list | - | GET | YES |
| Show Sprint | project/board/sprint/show/:sprint_id | - | GET | YES |
| Create Issue | project/board/:board_id/issue/create | JSON{ name, description, user_id?, estimate?, status?, sprint_id? } | POST | YES |
| Update Issue | project/board/issue/update/:issue_id | JSON{ name, description, user_id?, estimate?, status?, sprint_id? ...fields } | PUT | YES |
| Delete Issue | project/board/issue/delete/:issue_id | - | Delete | YES |
| Board Issues list | project/board/:board_id/issue/list | - | GET | YES |
| Sprint Issues list | project/board/sprint/:sprint_id/issue/list | - | GET | YES |
| Show Issue | project/board/issue/show/:issue_id | - | GET | YES |

