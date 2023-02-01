const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance)

    document.getElementById("go").addEventListener("click",() =>{
        document.body.innerHTML += SaySomething()
    })
})