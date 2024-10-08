# Use Vecty UI library to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

This demo is generated using [vectytemplater](https://github.com/soypat/vectytemplater) and slightly modified to make a HTTP request
by adding a button, and additional `actions.HttpRequest` Action.

Tested with Go 1.23.2, Vecty [v0.6.0](https://github.com/hexops/vecty/releases/tag/v0.6.0), 
vectytemplater [v0.0.0-20220501050640-d40b24e35168](https://pkg.go.dev/github.com/soypat/vectytemplater@v0.0.0-20220501050640-d40b24e35168).

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-vecty
```

2. Install project dependencies listed in `go.mod`:

```sh
go get
```

3. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o vecty.wasm main.go
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

2. As HTML, JS files, and a 10M-sized `vecty.wasm` are loaded into browser, press a "HTTP Request" button and refer to browser developer console
   to see the results.

### Finish

Perform your own experiments if desired.
