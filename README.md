# go-api-test

## SETUP

- Install [Cockroach DB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb.html)

- Start database
```
cockroach start --insecure
```

- Create an user
```
cockroach user set myuser --insecure
```

- Create database
```
cockroach sql --insecure -e 'CREATE DATABASE wwg'
```

- And create table(s)
```
create table musics (id serial primary key, artist string, title string);
```

## RUN
`go run ./main.go` :poop:
