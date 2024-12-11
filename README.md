# To Do List App
To Do List App

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

Test
go test -v ./...