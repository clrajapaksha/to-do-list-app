# To Do List App
To Do List App



### Generating docs
```
export PATH=$(go env GOPATH)/bin:$PATH
swag init -g cmd/main.go
```

Test
go test -v ./...