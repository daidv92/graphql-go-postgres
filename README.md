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

Query createMember & createSkill
```
mutation createMember {
  createMember(input:{name:"DaiDV"}) {
    id
    name
  }
}

mutation createSkill {
  createSkill(input:{category:"IT", name:"PHP", exp:7, memberId: 1}) {
    category
    name
    exp
  }
}
```
Query to get the all Members
```
query findMember {
  members(input: {}) {
    id
    name
    skill {
       category
       name
       exp
    }
  }
}
```
Query to get the all Skills
```
query findSkill {
  skills(input: {}){
    category
    name
    exp
  }
}
```
Query to get the all Members & Skills
```
query findAll{
  members(input: {ids: [1,2]}) {
    id
    name
    skill {
      category
      name
      exp
    }
  }
  skills(input: {ids: [1,2,3,4,5]}) {
    id
    category
    name
    exp
  }
}
```