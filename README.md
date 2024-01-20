# Go wasm demo

## setup 
1. set up go environment
2. settings must include GOOS = js GOARCH = WASM
3. go build -o helloworld.wasm helloworld.go
4. wasm_exec.js comes from the goroot/misc/ folder.

## demo
1. open wasmgo.html in browser
2. check out console should see results 
