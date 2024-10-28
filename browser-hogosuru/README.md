# Use Hogosuru framework to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

This demo is based on code from official Hogosuru examples like [learn-hogosuru](https://github.com/realPy/learn-hogosuru).

Tested with Go 1.23.3, Hogosuru [v1.5.1](https://github.com/realPy/hogosuru/releases/tag/v1.5.1).

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-hogosuru
```

2. Install project dependencies listed in `go.mod`:

```sh
go get
```

3. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

4. Copy the glue JS from Golang distribution to example's folder:

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
python3 -m http.server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As HTML, JS files, and a 3.2M-sized WASM file are loaded into browser, press a "HTTP Request" button and refer to browser developer console
   to see the results.

### Finish

Perform your own experiments if desired.
