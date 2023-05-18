- Set environment variable to 12 Factor App complient.

```shell
export HTTP_URL=http://0.0.0.0:8081
export WS_URL=wss://nostr.wine
export GRPC_URL=0.0.0.0:8080
export MONGODB_URL=mongodb://admin:password@mongodb:27017
```

- Get type-safe address in your application code.

```go
var (
	HTTP_URL    = env.HttpAddr("HTTP_URL")
	RELAY_URL   = env.WebsocketAddr("RELAY_URL")
	GRPC_URL    = env.GrpcAddr("GRPC_URL")
	MONGODB_URL = env.MongoAddr("MONGODB_URL")
)
```
