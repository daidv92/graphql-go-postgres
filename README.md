#Installation
Install the dependencies

go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
go get github.com/lib/pq

#pgadmin
Browse to localhost:5050 > Click add new server > Enter any name > Click on connection tab > Enter "database" for Hostname/Address > Enter "postgres-dev" for username > Enter "password" for password > Click save