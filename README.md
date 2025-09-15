# Coffee Shop gRPC Example

This is a simple Go project demonstrating a menu ordering system using gRPC. It includes both a server and a client, and uses Protocol Buffers for message definitions.

## Features
- View a menu of items from a coffee shop
- Place an order for selected items
- Check the status of your order

## Prerequisites
- [Go](https://golang.org/dl/) (1.18 or newer recommended)
- [protoc (Protocol Buffers Compiler)](https://grpc.io/docs/protoc-installation/)
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go) and [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc) plugins installed

## Getting Started

1. **Clone the repository**
   ```sh
   git clone git@github.com:kevinmulugu/golang-grpc-beginner-coffee-shop.git go-proto-example
   cd go-proto-example
   ```

2. **Generate gRPC code**
   Run the following command to generate Go code from the proto file:
   ```sh
   make
   ```
   This will create the generated files in the `coffeshop_protos/` directory.

3. **Run the server**
   In one terminal, start the gRPC server:
   ```sh
   go run server.go
   ```
   You should see:
   ```
   Coffee shop gRPC server is listening on :9001
   ```

4. **Run the client**
   In another terminal, run the client:
   ```sh
   go run client/client.go
   ```
   The client will:
   - Fetch the menu from the server
   - Place an order for all menu items
   - Check the status of the order

## File Structure
- `coffee_shop.proto` — Protocol Buffers definition
- `server.go` — gRPC server implementation
- `client/client.go` — gRPC client implementation
- `coffeshop_protos/` — Generated Go code from the proto file
- `Makefile` — Automates code generation

## Notes
- This project is for learning and demonstration purposes.
- The menu and order logic are hardcoded for simplicity.
- Make sure the server is running before starting the client.

## Troubleshooting
- If you get errors about missing generated files, make sure you ran `make` before running the server or client.
- If you change `coffee_shop.proto`, re-run `make` to regenerate the code.

## License

This project is licensed under the [MIT License](LICENSE).
