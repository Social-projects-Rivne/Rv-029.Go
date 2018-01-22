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