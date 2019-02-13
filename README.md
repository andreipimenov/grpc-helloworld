## GRPC Hello world

Start learning GRPC by writing hello world service which says hello to the user by his name.

### Build pb package from proto file

```
protoc --go_out=plugins=grpc:. pb/helloworld.proto
```

### Build server, client and run
```
cd cmd/server
go build
./server
```
```
2019/02/13 20:35:57 Listen on :8080
```
```
cd cmd/client
go build 
./client "John Doe"
```
```
2019/02/13 20:36:40 Response: Hello, John Doe
```
