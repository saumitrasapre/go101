

# to build:

First install protoc and the protoc-gen-go tools as per the gRPC docs

$ protoc --go_out=. --go_opt=paths=source_relative weblog/weblog.proto

# to run the console logging webserver:

$ go run cmd/loghttpd/main.go

# to run the protocol buffers webserver:

$ go run cmd/pbhttpd/main.go

# to run the log printer:

$ go run cmd/logprinter/main.go
