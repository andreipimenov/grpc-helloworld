## GRPC Hello world

Start learning GRPC by writing hello world service which says hello to the user by his name.

### Build pb package from proto file

```
protoc --go_out=plugins=grpc:. pb/helloworld.proto
```