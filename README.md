# learn-go
Repo to try out and learn go

## Modules
- `brotli` (in-progress)
- `protobuf`: A basic client server setup to try `protobuf` with `gRPC`

### protobuf

Steps to try this on your machine:
1. [Install](https://github.com/golang/protobuf#installation) `protobuf` binary on your machine
2. Run `make configure-protobuf` to generate protobuf related files
3. Run the server using `go run src/github.com/addityasingh/protobuf/server/main.go`
4. Run the client to verify the echoed response using `src/github.com/addityasingh/protobuf/client/main.go`

This should give output as below:
```bash
learn-go [master●●●] % go run src/github.com/addityasingh/protobuf/client/main.go
message:"Hello World!"
```