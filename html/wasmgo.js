const go = new Go();
go.importObject.env = {
}
WebAssembly.instantiateStreaming(fetch('helloworld.wasm'), go.importObject).then((result) => {
    go.run(result.instance);
    console.log("Result:", goAdd(2,3));
    console.log("Result:", goAdd(5,3));
    console.log("Result:", goAdd(12,3));
   });

