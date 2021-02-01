#Installation
Install app

> $ docker-compose build\
$ docker-compose up -d\
$ docker-compose start

Install the dependencies

> $ docker-compose exec server bash\
$ go get github.com/graphql-go/graphql\
$ go get github.com/graphql-go/handler\
$ go get github.com/lib/pq

#pgadmin
1. Browse to http://localhost:5050/browser/
2. Click add new server
3. Enter any name
4. Click on connection tab
5. Enter "database" for Hostname/Address
6. Enter "dev" for database name
7. Enter "postgres-dev" for username
8. Enter "password" for password
9. Click save

# Start the GraphQL server
> $ docker-compose exec server bash\
$ go run server.go

# Test
http://localhost:3004/

Query to get the all Members
```
query members {
  members{
      id
      name
  }
}
```
Query to get the all Skills
```
query skills {
  skills{
      category
      name
      exp
  }
}
```