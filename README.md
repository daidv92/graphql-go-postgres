#Installation
Install app

> $docker-compose build\
$docker-compose up -d\
$docker-compose start

Install the dependencies

> $docker-compose exec server bash\
$go get github.com/graphql-go/graphql\
$go get github.com/graphql-go/handler\
$go get github.com/lib/pq

#pgadmin
Browse to localhost:5050 > Click add new server > Enter any name > Click on connection tab > Enter "database" for Hostname/Address > Enter "postgres-dev" for username > Enter "password" for password > Click save

# Start it
go run server.go

# Test
Members
```
query members {
  members{
      id
      name
  }
}
```
Skills
```
query skills {
  skills{
      category
      name
      exp
  }
}
```