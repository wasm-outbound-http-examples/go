# Use stealthrocket/net library to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

2. Since [`stealthrocket/net`](https://github.com/stealthrocket/net) uses `GOOS=wasip1` feature to be
   [released](https://github.com/stealthrocket/net/blob/v0.1.4/README.md?plain=1#L18) 
   in Go 1.21 in August 2023, the installation of `gotip` (Golang from master branch) is needed prior that release.
   To install Gotip enter the command:

```sh
go install golang.org/dl/gotip@latest
gotip download
```

### Building

1. `cd` into the folder of this example:

```sh
cd wasi-stealthrocket-net
```

2. Install the `stealthrocket/net` library:

```sh
go get github.com/stealthrocket/net@v0.1.4
```

3. Compile examples using `gotip`:

```sh
GOOS=wasip1 GOARCH=wasm gotip build -o httpget.wasm httpget.go
GOOS=wasip1 GOARCH=wasm gotip build -o httpsget.wasm httpsget.go
```

### Test with wasirun

Wasirun is a CLI wrapper for Wazero runtime.

1. Install `wasirun` command. 

```sh
go install github.com/stealthrocket/wasi-go/cmd/wasirun@v0.5.0
```

2. Run WASM modules using `wasirun`:

```sh
wasirun httpget.wasm
wasirun httpsget.wasm
```

### Test with WasmEdge

WasmEdge runtime features its own implementation of WASI sockets.

Proven to work in WasmEdge 0.13 .

1. Install WasmEdge runtime:

```sh
curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash
```

2. Run WASM modules using `wasmedge`:

```sh
wasmedge httpget.wasm
wasmedge httpsget.wasm
```

### Finish

Perform your own experiments if desired.
