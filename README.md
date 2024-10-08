# Make HTTP requests from inside WASM in Golang

This devcontainer is configured to provide you a latest stable version of Go toolset.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/wasm-outbound-http-examples/go)

### WASI / Standalone Examples

<table>
<tr>
<th>#</th>
<th>Example</th>
<th>Description</th>
</tr>
<tr>
<td>1</td>
<td>

[stealthrocket/net](wasi-stealthrocket-net/README.md)

</td>
<td>

Use stealthrocket/net library to send HTTP requests from WASI envinronment.
Test resulting WASM with `wasirun` and `wasmedge`.

</td>
</tr>
</table>

### Browser / JS runtime Examples

<table>
<tr>
<th>#</th>
<th>Example</th>
<th>Description</th>
<th>Browser demo</th>
</tr>
<tr>
<td>1</td>
<td>

[Nigel2392/requester](browser-requester/README.md)

</td>
<td>

Use Nigel2392/requester library to send HTTP requests from web browser.

</td>
<td>

[Demo](https://wasm-outbound-http-examples.github.io/go/requester/)

</td>
</tr>
<tr>
<td>2</td>
<td>

[Vecty](browser-vecty/README.md)

</td>
<td>

Use Vecty UI library to send HTTP requests from web browser.

</td>
<td>

[Demo](https://wasm-outbound-http-examples.github.io/go/vecty/)

</td>
</tr>
</table>

<sub>Created for (wannabe-awesome) [list](https://github.com/vasilev/HTTP-request-from-inside-WASM)</sub>
