# Go wasm demo

## setup 
1. set up go environment
2. settings must include GOOS = js GOARCH = WASM
3. go build -o helloworld.wasm helloworld.go
4. wasm_exec.js comes from the goroot/misc/ folder.

## demo
1. docker build -f docker/Dockerfile -t my-wasm-demo .
2. docker run -dit --name my-running-app -p 8080:80 my-wasm-demo:latest
3. open localhost:8080/wasmgo.html in browser
4. check out console should see results 
