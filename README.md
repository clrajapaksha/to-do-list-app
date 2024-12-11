# To Do List App
To Do List App

* Swagger UI endpoint: http://localhost:8080/swagger/index.html

### Run server
```
go run cmd/main.go
```

### Run tests
```
go test -v ./...
```

### Generating docs
```
export PATH=$(go env GOPATH)/bin:$PATH
swag init -g cmd/main.go
```
